package csvparser

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Holding struct {
	Instrument   string
	Quantity     int
	AverageCost  float64
	LTP          float64
	Invested     float64
	CurrentValue float64
	PnL          float64
}

func ParseHoldingsCSV(filePath string) ([]Holding, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var holdings []Holding
	for i, row := range records {
		if i == 0 {
			continue // skip header
		}
		// Skip govt securities or empty lines
		if strings.Contains(row[0], "GS") || len(row[0]) == 0 {
			continue
		}

		qty, _ := strconv.Atoi(row[1])
		avgCost, _ := strconv.ParseFloat(row[2], 64)
		ltp, _ := strconv.ParseFloat(row[3], 64)
		invested, _ := strconv.ParseFloat(row[4], 64)
		curVal, _ := strconv.ParseFloat(row[5], 64)
		pnl, _ := strconv.ParseFloat(row[6], 64)

		holdings = append(holdings, Holding{
			Instrument:   row[0],
			Quantity:     qty,
			AverageCost:  avgCost,
			LTP:          ltp,
			Invested:     invested,
			CurrentValue: curVal,
			PnL:          pnl,
		})
	}

	return holdings, nil
}
