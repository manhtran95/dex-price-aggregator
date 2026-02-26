package dex

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/contracts"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

// GetTokenInfo fetches ERC-20 token metadata (symbol, name, decimals) from the chain.
// Shared by all DEX implementations to avoid duplication.
func GetTokenInfo(client *ethclient.Client, tokenAddress common.Address) (models.Token, error) {
	erc20, err := contracts.NewERC20(tokenAddress, client)
	if err != nil {
		return models.Token{}, err
	}

	symbol, err := erc20.Symbol(&bind.CallOpts{})
	if err != nil {
		return models.Token{}, err
	}

	name, err := erc20.Name(&bind.CallOpts{})
	if err != nil {
		return models.Token{}, err
	}

	decimals, err := erc20.Decimals(&bind.CallOpts{})
	if err != nil {
		return models.Token{}, err
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
