package dex

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/cache"
	"github.com/manhtran95/dex-price-aggregator/internal/contracts"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

// GetTokenInfo fetches ERC-20 token metadata (symbol, name, decimals).
// Results are cached in Redis (24-hour TTL) when a cache is provided.
// Safe to call with a nil cache — falls back to a direct RPC call.
func GetTokenInfo(ctx context.Context, client *ethclient.Client, c *cache.RedisCache, tokenAddress common.Address) (models.Token, error) {
	// --- cache read ---
	if c != nil {
		if meta, err := c.GetTokenMetadata(ctx, tokenAddress); err == nil && meta != nil {
			return models.Token{
				Address:  tokenAddress,
				Symbol:   meta.Symbol,
				Name:     meta.Name,
				Decimals: meta.Decimals,
			}, nil
		}
	}

	// --- RPC fetch ---
	erc20, err := contracts.NewERC20(tokenAddress, client)
	if err != nil {
		return models.Token{}, err
	}

	opts := &bind.CallOpts{Context: ctx}

	symbol, err := erc20.Symbol(opts)
	if err != nil {
		return models.Token{}, err
	}

	name, err := erc20.Name(opts)
	if err != nil {
		return models.Token{}, err
	}

	decimals, err := erc20.Decimals(opts)
	if err != nil {
		return models.Token{}, err
	}

	// --- cache write (non-blocking, errors silently ignored) ---
	if c != nil {
		_ = c.SetTokenMetadata(ctx, tokenAddress, symbol, name, decimals)
	}

	return models.Token{
		Address:  tokenAddress,
		Symbol:   symbol,
		Name:     name,
		Decimals: decimals,
	}, nil
}

// calculatePrice computes the human-readable output/input price,
// normalising both amounts by their respective token decimals.
// Shared by all DEX implementations.
func calculatePrice(amountOut, amountIn *big.Int, decimalsOut, decimalsIn uint8) float64 {
	outFloat := new(big.Float).SetInt(amountOut)
	inFloat := new(big.Float).SetInt(amountIn)

	outFloat.Quo(outFloat, big.NewFloat(math.Pow10(int(decimalsOut))))
	inFloat.Quo(inFloat, big.NewFloat(math.Pow10(int(decimalsIn))))

	price := new(big.Float).Quo(outFloat, inFloat)
	result, _ := price.Float64()
	return result
}
