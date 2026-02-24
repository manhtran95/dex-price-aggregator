package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Token represents an ERC20 token
type Token struct {
	Address  common.Address `json:"address"`
	Symbol   string         `json:"symbol"`
	Decimals uint8          `json:"decimals"`
	Name     string         `json:"name"`
}

// Quote represents a price quote from a DEX
type Quote struct {
	DEX          string   `json:"dex"`
	InputToken   Token    `json:"inputToken"`
	OutputToken  Token    `json:"outputToken"`
	InputAmount  *big.Int `json:"inputAmount"`
	OutputAmount *big.Int `json:"outputAmount"`
	Price        float64  `json:"price"`
	GasEstimate  uint64   `json:"gasEstimate,omitempty"`
}

type RouteResponse struct {
	BestRoute Route   `json:"bestRoute"`
	AllRoutes []Route `json:"allRoutes"`
}

type RouteWithGas struct {
	Route          *Route
	GasEstimate    uint64
	GasPriceGwei   float64
	GasCostETH     float64
	GasCostUSD     float64
	OutputValueUSD float64
	NetValueUSD    float64
}

// Route represents a swap route (can be multi-hop, multi-DEX)
type Route struct {
	Path        []Token  `json:"path"`        // [WETH, DAI, USDC]
	Hops        []Hop    `json:"hops"`        // Details of each swap
	TotalOutput *big.Int `json:"totalOutput"` // Final amount out
}

// Hop represents a single swap in a route
type Hop struct {
	DEX          string   `json:"dex"`           // "UniswapV3"
	TokenIn      Token    `json:"tokenIn"`       // WETH
	TokenOut     Token    `json:"tokenOut"`      // DAI
	InputAmount  *big.Int `json:"inputAmount"`   // 1 ETH
	OutputAmount *big.Int `json:"outputAmount"`  // 2000 DAI
	Price        float64  `json:"price"`         // 2000.00
	Fee          string   `json:"fee,omitempty"` // "0.3%" (for V3)
}

// SwapRequest represents an API request for a swap quote
type SwapRequest struct {
	InputToken  string `json:"inputToken"`
	OutputToken string `json:"outputToken"`
	Amount      string `json:"amount"`
}

// SwapResponse represents the API response
type SwapResponse struct {
	BestQuote Quote   `json:"bestQuote"`
	AllQuotes []Quote `json:"allQuotes"`
}
