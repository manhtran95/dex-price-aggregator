package dex

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

// MockDEX implements the DEX interface for testing
type MockDEX struct {
	NameFunc           func() string
	GetQuoteFunc       func(ctx context.Context, tokenIn, tokenOut common.Address, amountIn *big.Int) (*models.Quote, error)
	GetPairAddressFunc func(tokenA, tokenB common.Address) (common.Address, error)
}

func (m *MockDEX) Name() string {
	if m.NameFunc != nil {
		return m.NameFunc()
	}
	return "MockDEX"
}

func (m *MockDEX) GetQuote(ctx context.Context, tokenIn, tokenOut common.Address, amountIn *big.Int) (*models.Quote, error) {
	if m.GetQuoteFunc != nil {
		return m.GetQuoteFunc(ctx, tokenIn, tokenOut, amountIn)
	}
	return &models.Quote{
		DEX:          "MockDEX",
		InputAmount:  amountIn,
		OutputAmount: big.NewInt(2000000000), // 2000 USDC
	}, nil
}

func (m *MockDEX) GetPairAddress(tokenA, tokenB common.Address) (common.Address, error) {
	if m.GetPairAddressFunc != nil {
		return m.GetPairAddressFunc(tokenA, tokenB)
	}
	return common.HexToAddress("0x1234567890123456789012345678901234567890"), nil
}
