package dex

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/cache"
	"github.com/manhtran95/dex-price-aggregator/internal/contracts"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

type Curve struct {
	client *ethclient.Client
	pools  map[string]*CurvePoolInfo
	cache  *cache.RedisCache
}

type CurvePoolInfo struct {
	address common.Address
	tokens  []common.Address
	name    string
}

func NewCurve(client *ethclient.Client, redisCache *cache.RedisCache) (*Curve, error) {
	return &Curve{
		client: client,
		cache:  redisCache,
		pools: map[string]*CurvePoolInfo{
			"3pool": {
				address: common.HexToAddress("0xbEbc44782C7dB0a1A60Cb6fe97d0b483032FF1C7"),
				tokens: []common.Address{
					common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"), // DAI
					common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // USDC
					common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"), // USDT
				},
				name: "3pool",
			},
			"frax": {
				address: common.HexToAddress("0xDcEF968d416a41Cdac0ED8702fAC8128A64241A2"),
				tokens: []common.Address{
					common.HexToAddress("0x853d955aCEf822Db058eb8505911ED77F175b99e"), // FRAX
					common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"), // USDC
				},
				name: "frax",
			},
		},
	}, nil
}

func (c *Curve) Name() string {
	return "Curve"
}

// GetQuoteForSpecificPool gets quote from a specific Curve pool with token indices
func (c *Curve) GetQuoteForSpecificPool(
	ctx context.Context,
	poolAddress common.Address,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
	indexI, indexJ *big.Int,
) (*models.Quote, error) {

	// Create pool contract instance
	pool, err := contracts.NewCurvePool(poolAddress, c.client)
	if err != nil {
		return nil, err
	}

	// Fetch token metadata (cached)
	tokenInInfo, err := GetTokenInfo(ctx, c.client, c.cache, tokenIn)
	if err != nil {
		return nil, fmt.Errorf("failed to get input token info: %w", err)
	}
	tokenOutInfo, err := GetTokenInfo(ctx, c.client, c.cache, tokenOut)
	if err != nil {
		return nil, fmt.Errorf("failed to get output token info: %w", err)
	}

	// Call get_dy(i, j, dx)
	amountOut, err := pool.GetDy(&bind.CallOpts{Context: ctx}, indexI, indexJ, amountIn)
	if err != nil {
		return nil, fmt.Errorf("get_dy failed: %w", err)
	}

	price := calculatePrice(amountOut, amountIn, tokenOutInfo.Decimals, tokenInInfo.Decimals)

	log.Printf("Curve quote for specific pool - tokenIn, tokenOut, amountOut, price: %s %s %s %f", tokenIn.Hex(), tokenOut.Hex(), amountOut.String(), price)

	return &models.Quote{
		DEX:          c.Name(),
		InputToken:   tokenInInfo,
		OutputToken:  tokenOutInfo,
		InputAmount:  amountIn,
		OutputAmount: amountOut,
		Price:        price,
	}, nil
}

// func (c *Curve) findTokenIndices(
// 	poolInfo *CurvePoolInfo,
// 	tokenIn, tokenOut common.Address,
// ) (*big.Int, *big.Int, bool) {

// 	indexI := big.NewInt(-1)
// 	indexJ := big.NewInt(-1)

// 	// Find token indices
// 	for i, token := range poolInfo.tokens {
// 		if token == tokenIn {
// 			indexI = big.NewInt(int64(i))
// 		}
// 		if token == tokenOut {
// 			indexJ = big.NewInt(int64(i))
// 		}
// 	}

// 	if indexI.Cmp(big.NewInt(-1)) == 0 || indexJ.Cmp(big.NewInt(-1)) == 0 {
// 		return big.NewInt(-1), big.NewInt(-1), false
// 	}

// 	return indexI, indexJ, true
// }

func (c *Curve) GetPairAddress(tokenA, tokenB common.Address) (common.Address, error) {
	// Curve doesn't have pairs, it has pools
	return common.Address{}, fmt.Errorf("Curve uses pools, not pairs")
}

// GetQuote finds best quote across all Curve pools
func (c *Curve) GetQuote(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Quote, error) {

	return nil, fmt.Errorf("Curve.GetQuote not implemented")

	// var bestQuote *models.Quote

	// // Try each pool
	// for _, poolInfo := range c.pools {
	// 	// Find token indices in this pool
	// 	indexI, indexJ, found := c.findTokenIndices(poolInfo, tokenIn, tokenOut)
	// 	if !found {
	// 		continue
	// 	}

	// 	// Get quote from this pool
	// 	quote, err := c.getQuoteFromPool(ctx, poolInfo.address, indexI, indexJ, amountIn)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	// Keep best quote
	// 	if bestQuote == nil || quote.OutputAmount.Cmp(bestQuote.OutputAmount) > 0 {
	// 		bestQuote = quote
	// 	}
	// }

	// if bestQuote == nil {
	// 	return nil, fmt.Errorf("no valid Curve pool found for %s -> %s", tokenIn.Hex(), tokenOut.Hex())
	// }

	// return bestQuote, nil
}
