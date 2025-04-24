package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kumaryash90/portfolio-analyzer/internal/csvparser"
	"github.com/kumaryash90/portfolio-analyzer/internal/summary"
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

	// for _, h := range holdings {
	// 	fmt.Printf("%+v\n", h)
	// }
	// println()

	enriched := csvparser.MergeHoldingsWithDetails(holdings)

	// for _, e := range enriched {
	// 	fmt.Printf("%+v\n", e)
	// }

	// println()

	prompt := summary.GeneratePrompt(enriched)
	// fmt.Println(prompt)

	markdownTable := summary.GenerateMarkdownTable(enriched)
	fullReport := "# Portfolio Overview\n\n" + markdownTable + "\n\n" + prompt

	err = summary.WritePromptToFile(fullReport, "portfolio-summary.md")
	if err != nil {
		log.Fatal("Failed to write prompt to file: ", err)
	}

	fmt.Println("Prompt saved to portfolio-summary.md")

	summaryText, err := summary.SendToGPT(fullReport)
	if err != nil {
		log.Fatal("GPT request failed:", err)
	}

	fmt.Println("GPT Response:\n", summaryText)
}
