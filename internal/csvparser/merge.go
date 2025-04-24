package csvparser

import (
	"log"
	"strings"

	"github.com/kumaryash90/portfolio-analyzer/internal/screener"
)

type EnrichedHolding struct {
	Symbol       string
	Quantity     int
	AvgPrice     float64
	ROCE         string
	PE           string
	ProfitGrowth string
	DebtToEquity string
}

func MergeHoldingsWithDetails(holdings []Holding) []EnrichedHolding {
	var enriched []EnrichedHolding

	for _, h := range holdings {
		if strings.Contains(h.Instrument, "GS") {
			continue
		}

		detail, err := screener.ScrapeFundamentals(strings.ToUpper(strings.TrimSpace(h.Instrument)))
		if err != nil {
			log.Printf("Failed to fetch data for %s: %v", h.Instrument, err)
			continue
		}

		enriched = append(enriched, EnrichedHolding{
			Symbol:       h.Instrument,
			Quantity:     h.Quantity,
			AvgPrice:     h.AverageCost,
			ROCE:         detail.ROCE,
			PE:           detail.PE,
			ProfitGrowth: detail.ProfitGrowth,
			DebtToEquity: detail.DebtToEquity,
		})
	}

	return enriched
}
