package csvparser

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Holding struct {
	Instrument      string
	Quantity        int
	AveragePrice    float64
	LastTradedPrice float64
	CurrentValue    float64
	ProfitLoss      float64
}

func ParseHoldingCSV(filePath string) ([]Holding, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	var holdings []Holding
	isHeader := true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if isHeader {
			isHeader = false
			continue
		}

		qty, _ := strconv.Atoi(record[1])
		avg, _ := strconv.ParseFloat(record[2], 64)
		ltp, _ := strconv.ParseFloat(record[3], 64)
		val, _ := strconv.ParseFloat(record[4], 64)
		pnl, _ := strconv.ParseFloat(record[5], 64)

		holding := Holding{
			Instrument:      record[0],
			Quantity:        qty,
			AveragePrice:    avg,
			LastTradedPrice: ltp,
			CurrentValue:    val,
			ProfitLoss:      pnl,
		}

		holdings = append(holdings, holding)
	}

	return holdings, nil
}
