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

func (h *Handlers) GetBestRoute(w http.ResponseWriter, r *http.Request) {
	var req models.SwapRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tokenIn := common.HexToAddress(req.InputToken)
	tokenOut := common.HexToAddress(req.OutputToken)

	amount := new(big.Int)
	if _, ok := amount.SetString(req.Amount, 10); !ok {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	bestRoute, allRoutes, err := h.aggregator.FindBestRouteWithGraph(r.Context(), tokenIn, tokenOut, amount)
	if err != nil {
		http.Error(w, "Failed to find routes", http.StatusInternalServerError)
		return
	}

	response := models.RouteResponse{
		BestRoute: *bestRoute,
		AllRoutes: allRoutes,
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
