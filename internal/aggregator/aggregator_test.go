package aggregator

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/manhtran95/dex-price-aggregator/internal/dex"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

func TestGetBestQuote(t *testing.T) {
	// Create mock DEXs
	mockDEX1 := &dex.MockDEX{
		NameFunc: func() string {
			return "UniswapV2"
		},
		GetQuoteFunc: func(ctx context.Context, tokenIn, tokenOut common.Address, amountIn *big.Int) (*models.Quote, error) {
			return &models.Quote{
				DEX:          "UniswapV2",
				InputAmount:  amountIn,
				OutputAmount: big.NewInt(1990000000), // 1990 USDC
				Price:        1990.0,
			}, nil
		},
	}

	mockDEX2 := &dex.MockDEX{
		NameFunc: func() string {
			return "SushiSwap"
		},
		GetQuoteFunc: func(ctx context.Context, tokenIn, tokenOut common.Address, amountIn *big.Int) (*models.Quote, error) {
			return &models.Quote{
				DEX:          "SushiSwap",
				InputAmount:  amountIn,
				OutputAmount: big.NewInt(2005000000), // 2005 USDC - BEST!
				Price:        2005.0,
			}, nil
		},
	}

	mockDEX3 := &dex.MockDEX{
		NameFunc: func() string {
			return "PancakeSwap"
		},
		GetQuoteFunc: func(ctx context.Context, tokenIn, tokenOut common.Address, amountIn *big.Int) (*models.Quote, error) {
			return &models.Quote{
				DEX:          "PancakeSwap",
				InputAmount:  amountIn,
				OutputAmount: big.NewInt(1985000000), // 1985 USDC
				Price:        1985.0,
			}, nil
		},
	}

	// Create aggregator with mocks
	agg := NewAggregator([]dex.DEX{mockDEX1, mockDEX2, mockDEX3}, nil)

	// Test
	tokenIn := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	tokenOut := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	amountIn := big.NewInt(1000000000000000000) // 1 ETH

	bestQuote, allQuotes, err := agg.GetBestQuote(context.Background(), tokenIn, tokenOut, amountIn)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, bestQuote)
	assert.Equal(t, "SushiSwap", bestQuote.DEX)
	assert.Equal(t, big.NewInt(2005000000), bestQuote.OutputAmount)
	assert.Len(t, allQuotes, 3)
}

func TestGetBestQuote_NoDEXs(t *testing.T) {
	agg := NewAggregator([]dex.DEX{}, nil)

	tokenIn := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	tokenOut := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	amountIn := big.NewInt(1000000000000000000)

	bestQuote, allQuotes, err := agg.GetBestQuote(context.Background(), tokenIn, tokenOut, amountIn)

	assert.Error(t, err)
	assert.Nil(t, bestQuote)
	assert.Nil(t, allQuotes)
}

func TestGetBestQuote_AllDEXsFail(t *testing.T) {
	mockDEX := &dex.MockDEX{
		GetQuoteFunc: func(ctx context.Context, tokenIn, tokenOut common.Address, amountIn *big.Int) (*models.Quote, error) {
			return nil, assert.AnError
		},
	}

	agg := NewAggregator([]dex.DEX{mockDEX}, nil)

	tokenIn := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	tokenOut := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	amountIn := big.NewInt(1000000000000000000)

	bestQuote, allQuotes, err := agg.GetBestQuote(context.Background(), tokenIn, tokenOut, amountIn)

	assert.Error(t, err)
	assert.Nil(t, bestQuote)
	assert.Nil(t, allQuotes)
}
