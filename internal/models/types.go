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

// Route represents a swap route (can be multi-hop)
type Route struct {
	Path        []Token  `json:"path"`
	Quotes      []Quote  `json:"quotes"`
	TotalOutput *big.Int `json:"totalOutput"`
	PriceImpact float64  `json:"priceImpact"`
	GasEstimate uint64   `json:"gasEstimate"`
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
