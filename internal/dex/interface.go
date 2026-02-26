package dex

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

// DEX is the interface that all DEX implementations must satisfy
type DEX interface {
	Name() string
	GetQuote(ctx context.Context, tokenIn, tokenOut common.Address, amountIn *big.Int) (*models.Quote, error)
	GetPairAddress(tokenA, tokenB common.Address) (common.Address, error)
}
