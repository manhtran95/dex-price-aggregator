package dex

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/contracts"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

type UniswapV2 struct {
	client  *ethclient.Client
	factory *contracts.UniswapV2Factory
}

type Reserves struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}

func NewUniswapV2(client *ethclient.Client) (*UniswapV2, error) {
	factoryAddr := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	factory, err := contracts.NewUniswapV2Factory(factoryAddr, client)
	if err != nil {
		return nil, err
	}

	return &UniswapV2{
		client:  client,
		factory: factory,
	}, nil
}

func (u *UniswapV2) Name() string {
	return "UniswapV2"
}

func (u *UniswapV2) GetQuote(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Quote, error) {
	pairAddress, err := u.GetPairAddress(tokenIn, tokenOut)
	if err != nil {
		return nil, err
	}

	if pairAddress == (common.Address{}) {
		return nil, fmt.Errorf("pair does not exist")
	}

	// Fetch token info
	tokenInInfo, err := GetTokenInfo(u.client, tokenIn)
	if err != nil {
		return nil, fmt.Errorf("failed to get input token info: %w", err)
	}

	tokenOutInfo, err := GetTokenInfo(u.client, tokenOut)
	if err != nil {
		return nil, fmt.Errorf("failed to get output token info: %w", err)
	}

	reserves, err := u.getReserves(ctx, pairAddress, tokenIn, tokenOut)
	if err != nil {
		return nil, err
	}

	log.Println("Reserves:", reserves.Reserve0.String(), reserves.Reserve1.String())

	amountOut := u.calculateOutput(amountIn, reserves.Reserve0, reserves.Reserve1)

	price := calculatePrice(amountOut, amountIn, tokenOutInfo.Decimals, tokenInInfo.Decimals)

	return &models.Quote{
		DEX:          u.Name(),
		InputToken:   tokenInInfo,
		OutputToken:  tokenOutInfo,
		InputAmount:  amountIn,
		OutputAmount: amountOut,
		Price:        price,
	}, nil
}

func (u *UniswapV2) GetPairAddress(tokenA, tokenB common.Address) (common.Address, error) {
	// Clean! Just like Solidity ✅
	return u.factory.GetPair(&bind.CallOpts{}, tokenA, tokenB)
}

func (u *UniswapV2) getReserves(
	ctx context.Context,
	pairAddress, tokenA, tokenB common.Address,
) (*Reserves, error) {
	// Create pair contract instance
	pair, err := contracts.NewUniswapV2Pair(pairAddress, u.client)
	if err != nil {
		return nil, err
	}

	// Call getReserves - clean! ✅
	reserves, err := pair.GetReserves(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	// Sort reserves based on token order
	token0, _ := sortTokens(tokenA, tokenB)

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

func (u *UniswapV2) calculateOutput(amountIn, reserveIn, reserveOut *big.Int) *big.Int {
	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Mul(reserveIn, big.NewInt(1000))
	denominator.Add(denominator, amountInWithFee)
	return new(big.Int).Div(numerator, denominator)
}

func sortTokens(tokenA, tokenB common.Address) (common.Address, common.Address) {
	if tokenA.Hex() < tokenB.Hex() {
		return tokenA, tokenB
	}
	return tokenB, tokenA
}
