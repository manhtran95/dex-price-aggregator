package api

import (
	"encoding/json"
	"net/http"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/manhtran95/dex-price-aggregator/internal/aggregator"
	"github.com/manhtran95/dex-price-aggregator/internal/blockchain"
	"github.com/manhtran95/dex-price-aggregator/internal/config"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

type Handlers struct {
	client     *blockchain.Client
	config     *config.Config
	aggregator *aggregator.Aggregator
}

func NewHandlers(client *blockchain.Client, cfg *config.Config, agg *aggregator.Aggregator) *Handlers {
	return &Handlers{
		client:     client,
		config:     cfg,
		aggregator: agg,
	}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	blockNumber, err := h.client.GetBlockNumber(r.Context())
	if err != nil {
		http.Error(w, "Failed to connect to blockchain", http.StatusServiceUnavailable)
		return
	}

	response := map[string]interface{}{
		"status":      "ok",
		"blockNumber": blockNumber.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handlers) GetRouteWithGas(w http.ResponseWriter, r *http.Request) {
	const MAX_HOPS = 3
	var req models.SwapRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	tokenIn := common.HexToAddress(req.InputToken)
	tokenOut := common.HexToAddress(req.OutputToken)
	amount := new(big.Int)
	amount.SetString(req.Amount, 10)

	// Find best route considering gas
	outputTokenPriceUSD := 1.0 // USDC = $1
	bestRoute, allRoutes, err := h.aggregator.FindBestRouteWithGas(
		r.Context(),
		tokenIn,
		tokenOut,
		amount,
		MAX_HOPS, // max hops
		outputTokenPriceUSD,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get current gas price for display
	gasEstimator := h.aggregator.GetGasEstimator()
	gasPriceGwei, _ := gasEstimator.GetCurrentGasPrice(r.Context())

	response := RouteResponseWithGas{
		BestRoute: bestRoute,
		AllRoutes: allRoutes,
		GasInfo: &GasInfo{
			CurrentGasPriceGwei: gasPriceGwei,
			ETHPriceUSD:         gasEstimator.GetETHPrice(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handlers) ComparePrices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ComparePrices endpoint - coming soon",
	})
}
