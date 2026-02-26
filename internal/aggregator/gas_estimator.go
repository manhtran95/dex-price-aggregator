package aggregator

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

// Gas costs constants
const (
	BaseGasUniswapV2 = 120000
	BaseGasUniswapV3 = 130000
	BaseGasSushiSwap = 120000
	BaseGasCurve     = 180000

	GasPerHop        = 50000
	TokenApprovalGas = 46000
)

// GasEstimator handles all gas-related calculations
type GasEstimator struct {
	client         *ethclient.Client
	ethPriceUSD    float64
	defaultGasGwei float64
}

func NewGasEstimator(client *ethclient.Client) *GasEstimator {
	return &GasEstimator{
		client:         client,
		ethPriceUSD:    2500.0,
		defaultGasGwei: 50.0,
	}
}

// EstimateRoute calculates gas and net value for a route
func (g *GasEstimator) EstimateRoute(
	route *models.Route,
	gasPriceGwei float64,
	outputTokenPriceUSD float64,
) *models.RouteWithGas {

	// Estimate gas units
	gasUnits := g.estimateGasStatic(route)

	// Calculate costs
	gasCostETH := float64(gasUnits) * gasPriceGwei / 1e9
	gasCostUSD := gasCostETH * g.ethPriceUSD

	// Calculate output value
	outputValueUSD := g.calculateOutputValueUSD(
		route.TotalOutput,
		outputTokenPriceUSD,
		route.Path[len(route.Path)-1].Decimals,
	)

	// Net value after gas
	netValueUSD := outputValueUSD - gasCostUSD

	return &models.RouteWithGas{
		Route:          route,
		GasEstimate:    gasUnits,
		GasPriceGwei:   gasPriceGwei,
		GasCostETH:     gasCostETH,
		GasCostUSD:     gasCostUSD,
		OutputValueUSD: outputValueUSD,
		NetValueUSD:    netValueUSD,
	}
}

// estimateGasStatic provides fast static gas estimation
func (g *GasEstimator) estimateGasStatic(route *models.Route) uint64 {
	if len(route.Hops) == 0 {
		return 0
	}

	totalGas := uint64(0)

	// Base gas for each hop
	for _, hop := range route.Hops {
		totalGas += g.getBaseGasForDEX(hop.DEX)
	}

	// Multi-hop overhead
	if len(route.Hops) > 1 {
		totalGas += GasPerHop * uint64(len(route.Hops)-1)
	}

	// Token approval (assume first time)
	totalGas += TokenApprovalGas

	return totalGas
}

func (g *GasEstimator) getBaseGasForDEX(dex string) uint64 {
	switch dex {
	case "UniswapV2":
		return BaseGasUniswapV2
	case "UniswapV3":
		return BaseGasUniswapV3
	case "SushiSwap":
		return BaseGasSushiSwap
	case "Curve":
		return BaseGasCurve
	default:
		return 150000
	}
}

// EstimateGasAccurate simulates transaction for accurate estimation
func (g *GasEstimator) EstimateGasAccurate(
	ctx context.Context,
	route *models.Route,
	userAddress common.Address,
) (uint64, error) {

	// TODO: Build actual transaction and simulate
	// This requires router contract integration
	return 0, fmt.Errorf("accurate gas estimation not yet implemented")

	/*
	   // Build transaction data
	   txData := buildSwapCalldata(route, userAddress)

	   // Estimate gas
	   gasEstimate, err := g.client.EstimateGas(ctx, ethereum.CallMsg{
	       From:  userAddress,
	       To:    &routerAddress,
	       Data:  txData,
	   })
	   if err != nil {
	       return 0, err
	   }

	   // Add 20% buffer
	   return uint64(float64(gasEstimate) * 1.2), nil
	*/
}

// GetCurrentGasPrice fetches current network gas price
func (g *GasEstimator) GetCurrentGasPrice(ctx context.Context) (float64, error) {
	if g.client == nil {
		return g.defaultGasGwei, fmt.Errorf("no RPC client configured")
	}
	gasPrice, err := g.client.SuggestGasPrice(ctx)
	if err != nil {
		return g.defaultGasGwei, err
	}

	gasPriceGwei := float64(gasPrice.Int64()) / 1e9
	return gasPriceGwei, nil
}

// SetETHPrice updates ETH price (should be called periodically)
func (g *GasEstimator) SetETHPrice(priceUSD float64) {
	g.ethPriceUSD = priceUSD
}

// GetETHPrice returns current ETH price
func (g *GasEstimator) GetETHPrice() float64 {
	return g.ethPriceUSD
}

func (g *GasEstimator) calculateOutputValueUSD(
	amount *big.Int,
	tokenPriceUSD float64,
	decimals uint8,
) float64 {
	// Convert big.Int to float with decimals
	amountFloat := new(big.Float).SetInt(amount)
	divisor := new(big.Float).SetFloat64(math.Pow10(int(decimals)))
	amountFloat.Quo(amountFloat, divisor)

	tokenAmount, _ := amountFloat.Float64()
	return tokenAmount * tokenPriceUSD
}
