package aggregator

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

// Common intermediate tokens
var (
	WETH_ADDRESS = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	DAI_ADDRESS  = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	USDC_ADDRESS = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	USDT_ADDRESS = common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7")
	WBTC_ADDRESS = common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599")
)

func (a *Aggregator) FindBestRoute(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Route, []models.Route, error) {

	allRoutes := []models.Route{}

	// 1. Try direct route
	directRoute, err := a.tryDirectRoute(ctx, tokenIn, tokenOut, amountIn)
	if err == nil && directRoute != nil {
		allRoutes = append(allRoutes, *directRoute)
	}

	// 2. Try multi-hop routes
	multiHopRoutes := a.tryMultiHopRoutes(ctx, tokenIn, tokenOut, amountIn)
	allRoutes = append(allRoutes, multiHopRoutes...)

	// 3. Find best route
	if len(allRoutes) == 0 {
		return nil, nil, errors.New("no valid routes found")
	}

	bestRoute := findBestRoute(allRoutes)

	return &bestRoute, allRoutes, nil
}

func (a *Aggregator) tryDirectRoute(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Route, error) {

	// GetBestQuote already finds best DEX for this pair
	bestQuote, _, err := a.GetBestQuote(ctx, tokenIn, tokenOut, amountIn)
	if err != nil {
		return nil, err
	}

	// Create route with single hop
	route := &models.Route{
		Path: []models.Token{
			bestQuote.InputToken,
			bestQuote.OutputToken,
		},
		Hops: []models.Hop{
			{
				DEX:          bestQuote.DEX, // ← DEX name stored here!
				TokenIn:      bestQuote.InputToken,
				TokenOut:     bestQuote.OutputToken,
				InputAmount:  bestQuote.InputAmount,
				OutputAmount: bestQuote.OutputAmount,
				Price:        bestQuote.Price,
			},
		},
		TotalOutput: bestQuote.OutputAmount,
	}

	return route, nil
}

func (a *Aggregator) tryMultiHopRoutes(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) []models.Route {

	intermediates := a.getIntermediateTokens(tokenIn, tokenOut)
	routes := []models.Route{}

	for _, intermediate := range intermediates {
		// First hop: tokenIn → intermediate
		// GetBestQuote finds best DEX for this hop
		quote1, _, err := a.GetBestQuote(ctx, tokenIn, intermediate, amountIn)
		if err != nil {
			continue
		}

		// Second hop: intermediate → tokenOut
		// GetBestQuote finds best DEX for this hop (might be different DEX!)
		quote2, _, err := a.GetBestQuote(ctx, intermediate, tokenOut, quote1.OutputAmount)
		if err != nil {
			continue
		}

		// Create route with two hops (potentially different DEXs!)
		route := models.Route{
			Path: []models.Token{
				quote1.InputToken,
				quote1.OutputToken, // intermediate
				quote2.OutputToken,
			},
			Hops: []models.Hop{
				{
					DEX:          quote1.DEX, // e.g., "UniswapV3"
					TokenIn:      quote1.InputToken,
					TokenOut:     quote1.OutputToken,
					InputAmount:  quote1.InputAmount,
					OutputAmount: quote1.OutputAmount,
					Price:        quote1.Price,
				},
				{
					DEX:          quote2.DEX, // e.g., "Curve" (different DEX!)
					TokenIn:      quote2.InputToken,
					TokenOut:     quote2.OutputToken,
					InputAmount:  quote2.InputAmount,
					OutputAmount: quote2.OutputAmount,
					Price:        quote2.Price,
				},
			},
			TotalOutput: quote2.OutputAmount,
		}

		routes = append(routes, route)
	}

	return routes
}

func (a *Aggregator) getIntermediateTokens(
	tokenIn, tokenOut common.Address,
) []common.Address {

	candidates := []common.Address{
		WETH_ADDRESS,
		DAI_ADDRESS,
		USDC_ADDRESS,
		USDT_ADDRESS,
		WBTC_ADDRESS,
	}

	intermediates := []common.Address{}
	for _, token := range candidates {
		if token != tokenIn && token != tokenOut {
			intermediates = append(intermediates, token)
		}
	}

	return intermediates
}

func findBestRoute(routes []models.Route) models.Route {
	best := routes[0]
	for _, route := range routes {
		if route.TotalOutput.Cmp(best.TotalOutput) > 0 {
			best = route
		}
	}
	return best
}
