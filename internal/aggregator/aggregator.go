package aggregator

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/manhtran95/dex-price-aggregator/internal/dex"
	"github.com/manhtran95/dex-price-aggregator/internal/models"
)

type Aggregator struct {
	dexes []dex.DEX
}

func NewAggregator(dexes []dex.DEX) *Aggregator {
	return &Aggregator{dexes: dexes}
}

func (a *Aggregator) GetBestQuote(
	ctx context.Context,
	tokenIn, tokenOut common.Address,
	amountIn *big.Int,
) (*models.Quote, []models.Quote, error) {

	quoteChan := make(chan *models.Quote, len(a.dexes))
	var wg sync.WaitGroup

	for _, d := range a.dexes {
		wg.Add(1)
		go func(d dex.DEX) {
			defer wg.Done()
			quote, err := d.GetQuote(ctx, tokenIn, tokenOut, amountIn)
			if err != nil {
				log.Printf("Error from %s: %v", d.Name(), err)
				return
			}
			quoteChan <- quote
		}(d)
	}

	wg.Wait()
	close(quoteChan)

	var allQuotes []models.Quote
	for quote := range quoteChan {
		allQuotes = append(allQuotes, *quote)
	}

	if len(allQuotes) == 0 {
		return nil, nil, fmt.Errorf("no quotes available")
	}

	bestQuote := findBest(allQuotes)
	return bestQuote, allQuotes, nil
}

func findBest(quotes []models.Quote) *models.Quote {
	best := &quotes[0]
	for i := range quotes {
		if quotes[i].OutputAmount.Cmp(best.OutputAmount) > 0 {
			best = &quotes[i]
		}
	}
	return best
}
