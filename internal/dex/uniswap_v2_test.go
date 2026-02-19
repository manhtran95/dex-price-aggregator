package dex

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestCalculateOutput(t *testing.T) {
	tests := []struct {
		name       string
		amountIn   *big.Int
		reserveIn  *big.Int
		reserveOut *big.Int
		expected   *big.Int
	}{
		{
			name:       "1 ETH for USDC with large reserves",
			amountIn:   big.NewInt(1000000000000000000),                          // 1 ETH
			reserveIn:  new(big.Int).Mul(big.NewInt(50000), big.NewInt(1e18)),    // 50,000 ETH
			reserveOut: new(big.Int).Mul(big.NewInt(100000000), big.NewInt(1e6)), // 100M USDC
			expected:   big.NewInt(1993960240),                                   // ~1994 USDC (with 0.3% fee)
		},
		{
			name:       "Small amount",
			amountIn:   big.NewInt(100000000000000000), // 0.1 ETH
			reserveIn:  new(big.Int).Mul(big.NewInt(50000), big.NewInt(1e18)),
			reserveOut: new(big.Int).Mul(big.NewInt(100000000), big.NewInt(1e6)),
			expected:   big.NewInt(199399602), // ~199.4 USDC
		},
	}

	u := &UniswapV2{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := u.calculateOutput(tt.amountIn, tt.reserveIn, tt.reserveOut)

			// Allow small difference due to rounding
			diff := new(big.Int).Sub(result, tt.expected)
			diff.Abs(diff)

			maxDiff := big.NewInt(100) // Allow 100 wei difference
			assert.True(t, diff.Cmp(maxDiff) <= 0,
				"Expected %s, got %s, diff %s", tt.expected, result, diff)
		})
	}
}

func TestSortTokens(t *testing.T) {
	tokenA := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2") // WETH
	tokenB := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // USDC

	token0, token1 := sortTokens(tokenA, tokenB)

	// USDC address is lower than WETH
	assert.Equal(t, tokenB, token0)
	assert.Equal(t, tokenA, token1)
}
