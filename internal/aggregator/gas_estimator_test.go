package aggregator

import (
	"math/big"
	"testing"

	"github.com/manhtran95/dex-price-aggregator/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGasEstimator_EstimateRoute_SingleHop(t *testing.T) {
	estimator := &GasEstimator{
		ethPriceUSD: 2500.0,
	}

	route := &models.Route{
		Hops: []models.Hop{
			{DEX: "UniswapV2"},
		},
		TotalOutput: big.NewInt(2000000000), // 2000 USDC
		Path: []models.Token{
			{Decimals: 18},
			{Decimals: 6},
		},
	}

	result := estimator.EstimateRoute(route, 50.0, 1.0)

	// Should estimate around 166k gas (120k base + 46k approval)
	assert.Greater(t, result.GasEstimate, uint64(160000))
	assert.Less(t, result.GasEstimate, uint64(180000))

	// Gas cost should be reasonable
	assert.Greater(t, result.GasCostUSD, 10.0)
	assert.Less(t, result.GasCostUSD, 30.0)

	// Net value should be positive
	assert.Greater(t, result.NetValueUSD, 1970.0)
}

func TestGasEstimator_EstimateRoute_MultiHop(t *testing.T) {
	estimator := &GasEstimator{
		ethPriceUSD: 2500.0,
	}

	route := &models.Route{
		Hops: []models.Hop{
			{DEX: "UniswapV3"},
			{DEX: "Curve"},
		},
		TotalOutput: big.NewInt(2010000000), // 2010 USDC
		Path: []models.Token{
			{Decimals: 18},
			{Decimals: 18},
			{Decimals: 6},
		},
	}

	result := estimator.EstimateRoute(route, 50.0, 1.0)

	// Multi-hop should cost more gas
	assert.Greater(t, result.GasEstimate, uint64(300000))

	// But might still have better net value if output is higher
	t.Logf("Multi-hop net value: $%.2f", result.NetValueUSD)
}

func TestGasEstimator_CompareRoutes(t *testing.T) {
	estimator := &GasEstimator{
		ethPriceUSD: 2500.0,
	}

	// Route 1: Direct, lower output, lower gas
	route1 := &models.Route{
		Hops:        []models.Hop{{DEX: "UniswapV2"}},
		TotalOutput: big.NewInt(1998000000), // 1998 USDC
		Path:        []models.Token{{Decimals: 18}, {Decimals: 6}},
	}

	// Route 2: Multi-hop, higher output, higher gas
	route2 := &models.Route{
		Hops:        []models.Hop{{DEX: "UniswapV3"}, {DEX: "Curve"}},
		TotalOutput: big.NewInt(2003000000), // 2003 USDC
		Path:        []models.Token{{Decimals: 18}, {Decimals: 18}, {Decimals: 6}},
	}

	result1 := estimator.EstimateRoute(route1, 50.0, 1.0)
	result2 := estimator.EstimateRoute(route2, 50.0, 1.0)

	t.Logf("Route 1 (direct): Output=$%.2f, Gas=$%.2f, Net=$%.2f",
		result1.OutputValueUSD, result1.GasCostUSD, result1.NetValueUSD)
	t.Logf("Route 2 (multi-hop): Output=$%.2f, Gas=$%.2f, Net=$%.2f",
		result2.OutputValueUSD, result2.GasCostUSD, result2.NetValueUSD)

	// Route 1 might actually be better after gas!
	if result1.NetValueUSD > result2.NetValueUSD {
		t.Log("✓ Direct route wins after considering gas")
	} else {
		t.Log("✓ Multi-hop route wins despite higher gas")
	}
}
