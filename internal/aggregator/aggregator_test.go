package aggregator

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/manhtran95/dex-price-aggregator/internal/dex"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
	"github.com/stretchr/testify/assert"
)

// ─── Fixtures ────────────────────────────────────────────────────────────────

var (
	testWETH = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	testUSDC = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	testDAI  = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")

	tokenWETH = models.Token{Address: testWETH, Symbol: "WETH", Decimals: 18, Name: "Wrapped Ether"}
	tokenUSDC = models.Token{Address: testUSDC, Symbol: "USDC", Decimals: 6, Name: "USD Coin"}
	tokenDAI  = models.Token{Address: testDAI, Symbol: "DAI", Decimals: 18, Name: "Dai Stablecoin"}
)

// ─── Helpers ─────────────────────────────────────────────────────────────────

// newAggregatorForTest builds an Aggregator with a controlled graph – no real RPC needed.
func newAggregatorForTest(dexes []dex.DEX, g *Graph) *Aggregator {
	dexMap := make(map[string]dex.DEX)
	for _, d := range dexes {
		dexMap[d.Name()] = d
	}
	return &Aggregator{
		dexes:        dexMap,
		graph:        g,
		gasEstimator: NewGasEstimator(nil, nil), // nil client+cache → falls back to defaultGasGwei
	}
}

// newTestGraph builds a minimal Graph from a slice of PoolInfo descriptors.
func newTestGraph(pools []PoolInfo) *Graph {
	g := &Graph{
		nodes: make(map[common.Address]*Node),
		edges: []*Edge{},
	}
	for _, p := range pools {
		g.AddPool(p)
	}
	return g
}

// makeQuote is a concise constructor for mock quote responses.
func makeQuote(dexName string, in, out models.Token, inAmt, outAmt *big.Int) *models.Quote {
	return &models.Quote{
		DEX:          dexName,
		InputToken:   in,
		OutputToken:  out,
		InputAmount:  new(big.Int).Set(inAmt),
		OutputAmount: new(big.Int).Set(outAmt),
	}
}

// ─── Tests ────────────────────────────────────────────────────────────────────

// TestFindBestRouteWithGas_DirectHop: single pool, one DEX, one hop.
func TestFindBestRouteWithGas_DirectHop(t *testing.T) {
	amountIn := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 1 WETH
	usdcOut := big.NewInt(2000_000000)                                // 2000 USDC (6 dec)

	mockV2 := &dex.MockDEX{
		NameFunc: func() string { return "UniswapV2" },
		GetQuoteFunc: func(_ context.Context, _, _ common.Address, amt *big.Int) (*models.Quote, error) {
			return makeQuote("UniswapV2", tokenWETH, tokenUSDC, amt, usdcOut), nil
		},
	}

	g := newTestGraph([]PoolInfo{
		{Address: common.HexToAddress("0x0001"), Token0: testWETH, Token1: testUSDC, DEX: "UniswapV2", Fee: 3000},
	})
	agg := newAggregatorForTest([]dex.DEX{mockV2}, g)

	best, all, err := agg.FindBestRouteWithGas(context.Background(), testWETH, testUSDC, amountIn, 3, 1.0)

	assert.NoError(t, err)
	assert.NotNil(t, best)
	assert.Len(t, all, 1)
	assert.Len(t, best.Route.Hops, 1)
	assert.Equal(t, "UniswapV2", best.Route.Hops[0].DEX)
	assert.Equal(t, usdcOut, best.Route.TotalOutput)
	// Gas should be populated (base + approval)
	assert.Greater(t, best.GasEstimate, uint64(0))
	assert.Greater(t, best.GasCostUSD, 0.0)
}

// TestFindBestRouteWithGas_BestNetValue: two DEXes, different outputs;
// the route with the higher net value (output USD − gas USD) wins.
func TestFindBestRouteWithGas_BestNetValue(t *testing.T) {
	amountIn := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	mockV2 := &dex.MockDEX{
		NameFunc: func() string { return "UniswapV2" },
		GetQuoteFunc: func(_ context.Context, _, _ common.Address, amt *big.Int) (*models.Quote, error) {
			return makeQuote("UniswapV2", tokenWETH, tokenUSDC, amt, big.NewInt(1990_000000)), nil
		},
	}
	mockSushi := &dex.MockDEX{
		NameFunc: func() string { return "SushiSwap" },
		GetQuoteFunc: func(_ context.Context, _, _ common.Address, amt *big.Int) (*models.Quote, error) {
			return makeQuote("SushiSwap", tokenWETH, tokenUSDC, amt, big.NewInt(2005_000000)), nil
		},
	}

	g := newTestGraph([]PoolInfo{
		{Address: common.HexToAddress("0x0001"), Token0: testWETH, Token1: testUSDC, DEX: "UniswapV2", Fee: 3000},
		{Address: common.HexToAddress("0x0002"), Token0: testWETH, Token1: testUSDC, DEX: "SushiSwap", Fee: 3000},
	})
	agg := newAggregatorForTest([]dex.DEX{mockV2, mockSushi}, g)

	best, all, err := agg.FindBestRouteWithGas(context.Background(), testWETH, testUSDC, amountIn, 3, 1.0)

	assert.NoError(t, err)
	assert.NotNil(t, best)
	assert.Len(t, all, 2)
	// SushiSwap gives 2005 USDC > 1990 USDC, same gas cost → higher net value
	assert.Equal(t, "SushiSwap", best.Route.Hops[0].DEX)
	assert.Equal(t, big.NewInt(2005_000000), best.Route.TotalOutput)
	assert.Greater(t, best.NetValueUSD, 0.0)
}

// TestFindBestRouteWithGas_MultiHop: two-hop path WETH → DAI → USDC.
func TestFindBestRouteWithGas_MultiHop(t *testing.T) {
	amountIn := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 1 WETH
	// 1 WETH → 2000 DAI (18 dec)
	daiOut := new(big.Int).Mul(big.NewInt(2000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	// 2000 DAI → 1995 USDC (6 dec)
	usdcOut := big.NewInt(1995_000000)

	mockV2 := &dex.MockDEX{
		NameFunc: func() string { return "UniswapV2" },
		GetQuoteFunc: func(_ context.Context, _, tokenOut common.Address, amt *big.Int) (*models.Quote, error) {
			switch tokenOut {
			case testDAI:
				return makeQuote("UniswapV2", tokenWETH, tokenDAI, amt, daiOut), nil
			case testUSDC:
				return makeQuote("UniswapV2", tokenDAI, tokenUSDC, amt, usdcOut), nil
			}
			return nil, fmt.Errorf("unexpected output token %s", tokenOut.Hex())
		},
	}

	g := newTestGraph([]PoolInfo{
		{Address: common.HexToAddress("0x0001"), Token0: testWETH, Token1: testDAI, DEX: "UniswapV2", Fee: 3000},
		{Address: common.HexToAddress("0x0002"), Token0: testDAI, Token1: testUSDC, DEX: "UniswapV2", Fee: 3000},
	})
	agg := newAggregatorForTest([]dex.DEX{mockV2}, g)

	best, all, err := agg.FindBestRouteWithGas(context.Background(), testWETH, testUSDC, amountIn, 3, 1.0)

	assert.NoError(t, err)
	assert.NotNil(t, best)
	assert.Len(t, all, 1)
	// Two hops
	assert.Len(t, best.Route.Hops, 2)
	assert.Equal(t, "UniswapV2", best.Route.Hops[0].DEX)
	assert.Equal(t, "UniswapV2", best.Route.Hops[1].DEX)
	// Final output is the USDC amount from hop 2
	assert.Equal(t, usdcOut, best.Route.TotalOutput)
	// Path tokens: [WETH, DAI, USDC]
	assert.Len(t, best.Route.Path, 3)
	assert.Equal(t, "WETH", best.Route.Path[0].Symbol)
	assert.Equal(t, "DAI", best.Route.Path[1].Symbol)
	assert.Equal(t, "USDC", best.Route.Path[2].Symbol)
}

// TestFindBestRouteWithGas_NoPaths: empty graph → "no paths found".
func TestFindBestRouteWithGas_NoPaths(t *testing.T) {
	agg := newAggregatorForTest([]dex.DEX{}, newTestGraph(nil))

	_, _, err := agg.FindBestRouteWithGas(
		context.Background(), testWETH, testUSDC,
		big.NewInt(1_000_000_000_000_000_000), 3, 1.0,
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no paths found")
}

// TestFindBestRouteWithGas_AllPathsFail: graph has a path but no DEX is registered
// → every hop fails → "all paths failed".
func TestFindBestRouteWithGas_AllPathsFail(t *testing.T) {
	g := newTestGraph([]PoolInfo{
		{Address: common.HexToAddress("0x0001"), Token0: testWETH, Token1: testUSDC, DEX: "UniswapV2", Fee: 3000},
	})
	// Intentionally register NO DEXes
	agg := newAggregatorForTest([]dex.DEX{}, g)

	_, _, err := agg.FindBestRouteWithGas(
		context.Background(), testWETH, testUSDC,
		big.NewInt(1_000_000_000_000_000_000), 3, 1.0,
	)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "all paths failed")
}

// TestFindBestRouteWithGas_GasImpact: a route with more hops pays more gas;
// if the extra output doesn't cover the extra gas, the simpler route wins.
func TestFindBestRouteWithGas_GasImpact(t *testing.T) {
	amountIn := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)

	// Direct hop: WETH → USDC (1990 USDC, 1 hop)
	directOut := big.NewInt(2001_000000)
	// Two-hop: WETH → DAI → USDC (2001 USDC before gas, but 2 hops cost more gas)
	daiOut := new(big.Int).Mul(big.NewInt(2000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	multiOut := big.NewInt(2001_000000)

	mockV2 := &dex.MockDEX{
		NameFunc: func() string { return "UniswapV2" },
		GetQuoteFunc: func(_ context.Context, tokenIn, tokenOut common.Address, amt *big.Int) (*models.Quote, error) {
			switch tokenOut {
			case testUSDC:
				// Could be either direct WETH→USDC or multi-hop DAI→USDC
				if tokenIn.Hex() == testDAI.Hex() {
					return makeQuote("UniswapV2", tokenDAI, tokenUSDC, amt, multiOut), nil
				}
				return makeQuote("UniswapV2", tokenWETH, tokenUSDC, amt, directOut), nil
			case testDAI:
				return makeQuote("UniswapV2", tokenWETH, tokenDAI, amt, daiOut), nil
			}
			return nil, fmt.Errorf("unexpected output token %s", tokenOut.Hex())
		},
	}

	g := newTestGraph([]PoolInfo{
		// Direct path
		{Address: common.HexToAddress("0x0001"), Token0: testWETH, Token1: testUSDC, DEX: "UniswapV2", Fee: 3000},
		// Multi-hop path via DAI
		{Address: common.HexToAddress("0x0002"), Token0: testWETH, Token1: testDAI, DEX: "UniswapV2", Fee: 3000},
		{Address: common.HexToAddress("0x0003"), Token0: testDAI, Token1: testUSDC, DEX: "UniswapV2", Fee: 3000},
	})

	// outputTokenPriceUSD = 1.0 (1 USDC = $1), ETH price in estimator = $2500
	agg := newAggregatorForTest([]dex.DEX{mockV2}, g)
	agg.gasEstimator.SetETHPrice(2500.0)

	best, all, err := agg.FindBestRouteWithGas(context.Background(), testWETH, testUSDC, amountIn, 3, 1.0)

	assert.NoError(t, err)
	assert.NotNil(t, best)
	assert.Len(t, all, 2)

	// one-hop route should win
	assert.Len(t, best.Route.Hops, 1)
	loser := all[0]
	if all[0].NetValueUSD == best.NetValueUSD {
		loser = all[1]
	}
	assert.Greater(t, best.NetValueUSD, loser.NetValueUSD)
}
