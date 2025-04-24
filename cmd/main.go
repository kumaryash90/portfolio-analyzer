package main

import (
	"fmt"
	"os"

	"github.com/kumaryash90/portfolio-analyzer/internal/csvparser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Require csv file path. E.g.: go run cmd/main.go <csv-path>")
		return
	}

	path := os.Args[1]
	holdings, err := csvparser.ParseHoldingsCSV(path)

	if err != nil {
		fmt.Println("Error parsing csv: ", err)
	}

	for _, h := range holdings {
		fmt.Printf("%+v\n", h)
	}
	println()

	enriched := csvparser.MergeHoldingsWithDetails(holdings)

	for _, e := range enriched {
		fmt.Printf("%+v\n", e)
	}
}
