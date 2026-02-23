package dex

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/contracts"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

type SushiSwap struct {
	client  *ethclient.Client
	factory *contracts.UniswapV2Factory // Reuse UniswapV2 contract bindings
}

func NewSushiSwap(client *ethclient.Client) (*SushiSwap, error) {
	// SushiSwap factory address on Ethereum mainnet
	factoryAddr := common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
	factory, err := contracts.NewUniswapV2Factory(factoryAddr, client)
	if err != nil {
		return nil, err
	}

	return &SushiSwap{
		client:  client,
		factory: factory,
	}, nil
}

func (s *SushiSwap) Name() string {
	return "SushiSwap"
}

func (s *SushiSwap) GetQuote(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Quote, error) {
	log.Println("Getting SushiSwap quote for:", tokenIn.Hex(), tokenOut.Hex(), amountIn.String())
	
	pairAddress, err := s.GetPairAddress(tokenIn, tokenOut)
	if err != nil {
		return nil, err
	}

	if pairAddress == (common.Address{}) {
		return nil, fmt.Errorf("pair does not exist")
	}

	// Fetch token info
	tokenInInfo, err := s.getTokenInfo(tokenIn)
	if err != nil {
		return nil, fmt.Errorf("failed to get input token info: %w", err)
	}

	tokenOutInfo, err := s.getTokenInfo(tokenOut)
	if err != nil {
		return nil, fmt.Errorf("failed to get output token info: %w", err)
	}

	reserves, err := s.getReserves(ctx, pairAddress, tokenIn, tokenOut)
	if err != nil {
		return nil, err
	}

	log.Println("SushiSwap Reserves:", reserves.Reserve0.String(), reserves.Reserve1.String())

	amountOut := s.calculateOutput(amountIn, reserves.Reserve0, reserves.Reserve1)

	price := calculatePriceSushi(amountOut, amountIn, tokenOutInfo.Decimals, tokenInInfo.Decimals)

	return &models.Quote{
		DEX:          s.Name(),
		InputToken:   tokenInInfo,
		OutputToken:  tokenOutInfo,
		InputAmount:  amountIn,
		OutputAmount: amountOut,
		Price:        price,
	}, nil
}

func (s *SushiSwap) GetPairAddress(tokenA, tokenB common.Address) (common.Address, error) {
	return s.factory.GetPair(&bind.CallOpts{}, tokenA, tokenB)
}

func (s *SushiSwap) getTokenInfo(tokenAddress common.Address) (models.Token, error) {
	erc20, err := contracts.NewERC20(tokenAddress, s.client)
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

func calculatePriceSushi(amountOut, amountIn *big.Int, decimalsOut, decimalsIn uint8) float64 {
	outFloat := new(big.Float).SetInt(amountOut)
	inFloat := new(big.Float).SetInt(amountIn)

	outFloat.Quo(outFloat, big.NewFloat(math.Pow10(int(decimalsOut))))
	inFloat.Quo(inFloat, big.NewFloat(math.Pow10(int(decimalsIn))))

	price := new(big.Float).Quo(outFloat, inFloat)
	result, _ := price.Float64()
	return result
}

func (s *SushiSwap) getReserves(
	ctx context.Context,
	pairAddress, tokenA, tokenB common.Address,
) (*Reserves, error) {
	// Create pair contract instance (reuse UniswapV2Pair binding)
	pair, err := contracts.NewUniswapV2Pair(pairAddress, s.client)
	if err != nil {
		return nil, err
	}

	// Call getReserves
	reserves, err := pair.GetReserves(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	// Sort reserves based on token order
	token0, _ := sortTokensSushi(tokenA, tokenB)

	if tokenA == token0 {
		return &Reserves{
			Reserve0: reserves.Reserve0,
			Reserve1: reserves.Reserve1,
		}, nil
	}

	return &Reserves{
		Reserve0: reserves.Reserve1,
		Reserve1: reserves.Reserve0,
	}, nil
}

func (s *SushiSwap) calculateOutput(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	// SushiSwap uses the same 0.3% fee as Uniswap V2
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Mul(reserveIn, big.NewInt(1000))
	denominator.Add(denominator, amountInWithFee)
	return new(big.Int).Div(numerator, denominator)
}

func sortTokensSushi(tokenA, tokenB common.Address) (common.Address, common.Address) {
	if tokenA.Hex() < tokenB.Hex() {
		return tokenA, tokenB
	}
	return tokenB, tokenA
}
