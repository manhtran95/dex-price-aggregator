package dex

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

const quoterV2ABI = `[{"inputs":[{"components":[{"internalType":"address","name":"tokenIn","type":"address"},{"internalType":"address","name":"tokenOut","type":"address"},{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint24","name":"fee","type":"uint24"},{"internalType":"uint160","name":"sqrtPriceLimitX96","type":"uint160"}],"internalType":"struct IQuoterV2.QuoteExactInputSingleParams","name":"params","type":"tuple"}],"name":"quoteExactInputSingle","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint160","name":"sqrtPriceX96After","type":"uint160"},{"internalType":"uint32","name":"initializedTicksCrossed","type":"uint32"},{"internalType":"uint256","name":"gasEstimate","type":"uint256"}],"stateMutability":"nonpayable","type":"function"}]`

type UniswapV3 struct {
	client      *ethclient.Client
	quoterAddr  common.Address
	quoterABI   abi.ABI
	factoryAddr common.Address
}

func NewUniswapV3(client *ethclient.Client) (*UniswapV3, error) {
	parsedABI, err := abi.JSON(strings.NewReader(quoterV2ABI))
	if err != nil {
		return nil, err
	}

	return &UniswapV3{
		client:      client,
		quoterAddr:  common.HexToAddress("0x61fFE014bA17989E743c5F6cB21bF9697530B21e"), // QuoterV2
		quoterABI:   parsedABI,
		factoryAddr: common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
	}, nil
}

func (u *UniswapV3) Name() string {
	return "UniswapV3"
}

func (u *UniswapV3) GetQuoteForSpecificPool(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
	fee uint32,
) (*models.Quote, error) {

	// Fetch token info
	tokenInInfo, err := GetTokenInfo(u.client, tokenIn)
	if err != nil {
		return nil, fmt.Errorf("failed to get input token info: %w", err)
	}

	tokenOutInfo, err := GetTokenInfo(u.client, tokenOut)
	if err != nil {
		return nil, fmt.Errorf("failed to get output token info: %w", err)
	}

	// Get quote for ONLY this fee tier
	amountOut, err := u.getQuoteForFee(ctx, tokenIn, tokenOut, amountIn, fee)
	if err != nil {
		return nil, err
	}

	price := calculatePrice(amountOut, amountIn, tokenOutInfo.Decimals, tokenInInfo.Decimals)

	log.Printf("UniswapV3 quote for specific pool - tokenIn, tokenOut, amountOut, fee: %s %s %s %d", tokenIn.Hex(), tokenOut.Hex(), amountOut.String(), fee)

	return &models.Quote{
		DEX:          u.Name(),
		InputToken:   tokenInInfo,
		OutputToken:  tokenOutInfo,
		InputAmount:  amountIn,
		OutputAmount: amountOut,
		Price:        price,
	}, nil
}

func (u *UniswapV3) GetQuote(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Quote, error) {

	// Fetch token info
	tokenInInfo, err := GetTokenInfo(u.client, tokenIn)
	if err != nil {
		return nil, fmt.Errorf("failed to get input token info: %w", err)
	}

	tokenOutInfo, err := GetTokenInfo(u.client, tokenOut)
	if err != nil {
		return nil, fmt.Errorf("failed to get output token info: %w", err)
	}

	// V3 has multiple fee tiers
	fees := []uint32{500, 3000, 10000} // 0.05%, 0.3%, 1%

	var bestQuote *models.Quote

	for _, fee := range fees {
		amountOut, err := u.getQuoteForFee(ctx, tokenIn, tokenOut, amountIn, fee)
		if err != nil {
			// Pool might not exist for this fee tier
			log.Printf("No pool found for fee %d", fee)
			log.Printf("Error: %v", err)
			continue
		}

		// Keep the best quote
		if bestQuote == nil || amountOut.Cmp(bestQuote.OutputAmount) > 0 {
			// Calculate price
			price := calculatePrice(amountOut, amountIn, tokenOutInfo.Decimals, tokenInInfo.Decimals)

			bestQuote = &models.Quote{
				DEX:          u.Name(),
				InputToken:   tokenInInfo,
				OutputToken:  tokenOutInfo,
				InputAmount:  amountIn,
				OutputAmount: amountOut,
				Price:        price,
			}
		}
	}

	if bestQuote == nil {
		return nil, fmt.Errorf("no valid V3 pools found")
	}

	return bestQuote, nil
}

func (u *UniswapV3) getQuoteForFee(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
	fee uint32,
) (*big.Int, error) {

	// Encode the struct as a tuple
	params := struct {
		TokenIn           common.Address
		TokenOut          common.Address
		AmountIn          *big.Int
		Fee               *big.Int // Must be *big.Int for ABI encoding
		SqrtPriceLimitX96 *big.Int
	}{
		TokenIn:           tokenIn,
		TokenOut:          tokenOut,
		AmountIn:          amountIn,
		Fee:               big.NewInt(int64(fee)),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	// Pack the function call
	data, err := u.quoterABI.Pack("quoteExactInputSingle", params)
	if err != nil {
		return nil, err
	}

	// Call the contract
	result, err := u.client.CallContract(ctx, ethereum.CallMsg{
		To:   &u.quoterAddr,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	// Unpack the result
	var output struct {
		AmountOut               *big.Int
		SqrtPriceX96After       *big.Int
		InitializedTicksCrossed uint32
		GasEstimate             *big.Int
	}

	err = u.quoterABI.UnpackIntoInterface(&output, "quoteExactInputSingle", result)
	if err != nil {
		return nil, err
	}

	return output.AmountOut, nil
}

// GetPairAddress is not used for V3, but included for interface compatibility
func (u *UniswapV3) GetPairAddress(tokenA, tokenB common.Address) (common.Address, error) {
	// V3 doesn't have a single pair per token pair
	// It has multiple pools with different fees
	return common.Address{}, fmt.Errorf("V3 uses pools, not pairs")
}
