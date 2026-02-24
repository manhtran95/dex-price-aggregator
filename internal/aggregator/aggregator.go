package aggregator

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/dex"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

type Aggregator struct {
	dexes        map[string]dex.DEX // Changed to map for lookup by name
	graph        *Graph
	gasEstimator *GasEstimator // ← Add this
	client       *ethclient.Client
}

func NewAggregator(dexes []dex.DEX, client *ethclient.Client) *Aggregator {
	dexMap := make(map[string]dex.DEX)
	for _, d := range dexes {
		dexMap[d.Name()] = d
	}

	return &Aggregator{
		dexes:        dexMap,
		graph:        NewGraph(),
		gasEstimator: NewGasEstimator(client), // ← Initialize
		client:       client,
	}
}

func (a *Aggregator) GetGasEstimator() *GasEstimator {
	return a.gasEstimator
}

// GetQuoteFromSpecificPool gets quote from a specific pool/DEX
func (a *Aggregator) GetQuoteFromSpecificPool(
	ctx context.Context,
	edge *Edge,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Quote, error) {

	// Get the specific DEX
	dexInstance, exists := a.dexes[edge.dex]
	if !exists {
		return nil, fmt.Errorf("DEX %s not found", edge.dex)
	}

	// For UniswapV3, we need to pass the specific fee tier
	if edge.dex == "UniswapV3" {
		// Call V3-specific method with fee
		v3, ok := dexInstance.(*dex.UniswapV3)
		if !ok {
			return nil, fmt.Errorf("failed to cast to UniswapV3")
		}
		return v3.GetQuoteForSpecificPool(ctx, tokenIn, tokenOut, amountIn, edge.fee)
	}

	// For V2 and others, just call GetQuote normally
	return dexInstance.GetQuote(ctx, tokenIn, tokenOut, amountIn)
}

// CalculatePathOutput calculates the output for a specific path
func (a *Aggregator) CalculatePathOutput(
	ctx context.Context,
	path Path,
	amountIn *big.Int,
) (*models.Route, error) {

	if len(path.Edges) == 0 {
		return nil, fmt.Errorf("path has no edges")
	}

	currentAmount := amountIn
	hops := []models.Hop{}
	// Collect enriched token metadata from quotes as we go
	enrichedTokens := []models.Token{}

	// Execute each hop in the path
	for i, edge := range path.Edges {
		tokenIn := path.Tokens[i]
		tokenOut := path.Tokens[i+1]

		// Get quote from the SPECIFIC pool in this edge
		quote, err := a.GetQuoteFromSpecificPool(ctx, edge, tokenIn, tokenOut, currentAmount)
		if err != nil {
			return nil, fmt.Errorf("hop %d failed: %w", i, err)
		}

		// Harvest token metadata from quote
		if i == 0 {
			enrichedTokens = append(enrichedTokens, quote.InputToken)
		}
		enrichedTokens = append(enrichedTokens, quote.OutputToken)

		// Add hop details
		hop := models.Hop{
			DEX:          edge.dex,
			TokenIn:      quote.InputToken,
			TokenOut:     quote.OutputToken,
			InputAmount:  currentAmount,
			OutputAmount: quote.OutputAmount,
			Price:        quote.Price,
			Fee:          fmt.Sprintf("%.2f%%", float64(edge.fee)/10000),
		}
		hops = append(hops, hop)

		// Output of this hop becomes input of next hop
		currentAmount = quote.OutputAmount
	}

	// Build final route
	route := &models.Route{
		Path:        enrichedTokens,
		Hops:        hops,
		TotalOutput: currentAmount,
	}

	return route, nil
}

func (a *Aggregator) FindBestRouteWithGas(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
	maxHops int,
	outputTokenPriceUSD float64,
) (*models.RouteWithGas, []*models.RouteWithGas, error) {

	// Get current gas price
	gasPriceGwei, err := a.gasEstimator.GetCurrentGasPrice(ctx)
	if err != nil {
		gasPriceGwei = 50.0 // Fallback
	}

	// Find all possible routes
	paths := a.graph.FindAllPaths(tokenIn, tokenOut, maxHops)
	if len(paths) == 0 {
		return nil, nil, fmt.Errorf("no paths found")
	}

	var bestRoute *models.RouteWithGas
	allRoutesWithGas := []*models.RouteWithGas{}

	for _, path := range paths {
		// Calculate route output
		route, err := a.CalculatePathOutput(ctx, path, amountIn)
		if err != nil {
			continue
		}

		// Estimate gas and calculate net value
		routeWithGas := a.gasEstimator.EstimateRoute(
			route,
			gasPriceGwei,
			outputTokenPriceUSD,
		)

		allRoutesWithGas = append(allRoutesWithGas, routeWithGas)

		// Keep route with best NET value (not just highest output)
		if bestRoute == nil || routeWithGas.NetValueUSD > bestRoute.NetValueUSD {
			bestRoute = routeWithGas
		}
	}

	if bestRoute == nil {
		return nil, nil, fmt.Errorf("all paths failed")
	}

	return bestRoute, allRoutesWithGas, nil
}

// 1st version: get best quote from all DEXs
func (a *Aggregator) GetBestQuote(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Quote, []models.Quote, error) {

	quoteChan := make(chan *models.Quote, len(a.dexes))
	var wg sync.WaitGroup

	for _, d := range a.dexes {
		wg.Add(1)
		go func(d dex.DEX) {
			defer wg.Done()
			quote, err := d.GetQuote(ctx, tokenIn, tokenOut, amountIn)
			if err != nil {
				log.Printf("Error from %s: %v", d.Name(), err)
				return
			}
			quoteChan <- quote
		}(d)
	}

	wg.Wait()
	close(quoteChan)

	var allQuotes []models.Quote
	for quote := range quoteChan {
		allQuotes = append(allQuotes, *quote)
	}

	if len(allQuotes) == 0 {
		return nil, nil, fmt.Errorf("no quotes available")
	}

	bestQuote := findBest(allQuotes)
	return bestQuote, allQuotes, nil
}

func findBest(quotes []models.Quote) *models.Quote {
	best := &quotes[0]
	for i := range quotes {
		if quotes[i].OutputAmount.Cmp(best.OutputAmount) > 0 {
			best = &quotes[i]
		}
	}
	return best
}
