package api

import "github.com/manhtran95/dex-price-aggregator/internal/models"

type RouteResponseWithGas struct {
	BestRoute *models.RouteWithGas   `json:"bestRoute"`
	AllRoutes []*models.RouteWithGas `json:"allRoutes"`
	GasInfo   *GasInfo        `json:"gasInfo"`
}

type GasInfo struct {
	CurrentGasPriceGwei float64 `json:"currentGasPriceGwei"`
	ETHPriceUSD         float64 `json:"ethPriceUSD"`
}
