package api

import (
	"encoding/json"
	"net/http"

	"github.com/manhtran95/dex-price-aggregator/internal/blockchain"
	"github.com/manhtran95/dex-price-aggregator/internal/config"
)

type Handlers struct {
	client *blockchain.Client
	config *config.Config
}

func NewHandlers(client *blockchain.Client, cfg *config.Config) *Handlers {
	return &Handlers{
		client: client,
		config: cfg,
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

func (h *Handlers) GetQuote(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement quote logic
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "GetQuote endpoint - coming soon",
	})
}

func (h *Handlers) ComparePrices(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement price comparison logic
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "ComparePrices endpoint - coming soon",
	})
}
