package integration

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/contracts"
	"github.com/manhtran95/dex-price-aggregator/internal/dex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// These tests require ETHEREUM_MAINNET_RPC environment variable
// Run with: ETHEREUM_MAINNET_RPC=https://mainnet.infura.io/v3/YOUR-KEY go test ./test/integration/...

func getClient(t *testing.T) *ethclient.Client {
	rpcURL := os.Getenv("ETHEREUM_MAINNET_RPC")
	if rpcURL == "" {
		t.Skip("ETHEREUM_MAINNET_RPC not set, skipping integration test")
	}

	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err, "Failed to connect to Ethereum")

	return client
}

func TestUniswapV2Factory_GetPair_RealMainnet(t *testing.T) {
	client := getClient(t)
	defer client.Close()

	// Uniswap V2 Factory on mainnet
	factoryAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	factory, err := contracts.NewUniswapV2Factory(factoryAddress, client)
	require.NoError(t, err)

	// WETH and USDC addresses
	weth := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")

	// Get pair address
	pairAddress, err := factory.GetPair(&bind.CallOpts{}, weth, usdc)
	require.NoError(t, err)

	// Expected WETH/USDC pair address
	expectedPair := common.HexToAddress("0xB4e16d0168e52d35CaCD2c6185b44281Ec28C9Dc")

	assert.Equal(t, expectedPair, pairAddress, "Should return correct WETH/USDC pair address")
	assert.NotEqual(t, common.Address{}, pairAddress, "Pair address should not be zero")

	t.Logf("✓ WETH/USDC Pair Address: %s", pairAddress.Hex())
}


func TestERC20_GetTokenInfo_WETH(t *testing.T) {
	client := getClient(t)
	defer client.Close()

	wethAddress := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	weth, err := contracts.NewERC20(wethAddress, client)
	require.NoError(t, err)

	// Get symbol
	symbol, err := weth.Symbol(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, "WETH", symbol)

	// Get name
	name, err := weth.Name(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, "Wrapped Ether", name)

	// Get decimals
	decimals, err := weth.Decimals(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, uint8(18), decimals)

	t.Logf("✓ WETH Token Info:")
	t.Logf("  Symbol: %s", symbol)
	t.Logf("  Name: %s", name)
	t.Logf("  Decimals: %d", decimals)
}

func TestERC20_GetTokenInfo_USDC(t *testing.T) {
	client := getClient(t)
	defer client.Close()

	usdcAddress := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	usdc, err := contracts.NewERC20(usdcAddress, client)
	require.NoError(t, err)

	symbol, err := usdc.Symbol(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, "USDC", symbol)

	name, err := usdc.Name(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, "USD Coin", name)

	decimals, err := usdc.Decimals(&bind.CallOpts{})
	require.NoError(t, err)
	assert.Equal(t, uint8(6), decimals)

	t.Logf("✓ USDC Token Info:")
	t.Logf("  Symbol: %s", symbol)
	t.Logf("  Name: %s", name)
	t.Logf("  Decimals: %d", decimals)
}

func TestUniswapV2_GetQuote_EndToEnd(t *testing.T) {
	client := getClient(t)
	defer client.Close()

	// Create UniswapV2 instance
	uniswapV2, err := dex.NewUniswapV2(client)
	require.NoError(t, err)

	// Test with WETH → USDC
	weth := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	usdc := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	amountIn := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e18)) // 1 ETH

	quote, err := uniswapV2.GetQuote(context.Background(), weth, usdc, amountIn)
	require.NoError(t, err)
	require.NotNil(t, quote)

	// Verify quote structure
	assert.Equal(t, "UniswapV2", quote.DEX)
	assert.Equal(t, weth, quote.InputToken.Address)
	assert.Equal(t, usdc, quote.OutputToken.Address)
	assert.Equal(t, "WETH", quote.InputToken.Symbol)
	assert.Equal(t, "USDC", quote.OutputToken.Symbol)
	assert.Equal(t, uint8(18), quote.InputToken.Decimals)
	assert.Equal(t, uint8(6), quote.OutputToken.Decimals)
	assert.Equal(t, amountIn, quote.InputAmount)

	// Verify output amount is reasonable (should be > 0 and < 10000 USDC for 1 ETH)
	minOutput := big.NewInt(100000000)   // 100 USDC (in 6 decimals)
	maxOutput := big.NewInt(10000000000) // 10,000 USDC (in 6 decimals)
	assert.True(t, quote.OutputAmount.Cmp(minOutput) > 0, "Output should be > 100 USDC")
	assert.True(t, quote.OutputAmount.Cmp(maxOutput) < 0, "Output should be < 10,000 USDC")

	// Verify price is reasonable (should be between 100 and 10000)
	assert.Greater(t, quote.Price, 100.0, "Price should be > 100")
	assert.Less(t, quote.Price, 10000.0, "Price should be < 10,000")

	t.Logf("✓ Quote for 1 WETH → USDC:")
	t.Logf("  Input: 1 WETH")
	t.Logf("  Output: %s USDC", new(big.Float).Quo(
		new(big.Float).SetInt(quote.OutputAmount),
		big.NewFloat(1e6),
	).Text('f', 2))
	t.Logf("  Price: %.2f USDC per WETH", quote.Price)
}
