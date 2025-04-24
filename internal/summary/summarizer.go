package summary

import (
	"fmt"
	"strings"

	"github.com/kumaryash90/portfolio-analyzer/internal/csvparser"
)

func GeneratePrompt(enriched []csvparser.EnrichedHolding) string {
	var b strings.Builder

	b.WriteString("You're a financial analyst. Analyze the following stock portfolio and provide insights.\n")
	b.WriteString("For each stock, highlight strengths, risks, and whether it's overvalued or undervalued.\n\n")

	for _, h := range enriched {
		b.WriteString(fmt.Sprintf("%s\n", h.Symbol))
		b.WriteString(fmt.Sprintf("  - Quantity: %d\n", h.Quantity))
		b.WriteString(fmt.Sprintf("  - Avg Buy Price: ₹%.2f\n", h.AvgPrice))

		if h.ROCE != "" {
			b.WriteString(fmt.Sprintf("  - ROCE: %s\n", h.ROCE))
		}
		if h.PE != "" {
			b.WriteString(fmt.Sprintf("  - PE Ratio: %s\n", h.PE))
		}
		if h.ProfitGrowth != "" {
			b.WriteString(fmt.Sprintf("  - Profit Growth (3Y): %s\n", h.ProfitGrowth))
		}
		if h.DebtToEquity != "" {
			b.WriteString(fmt.Sprintf("  - Debt to Equity: %s\n", h.DebtToEquity))
		}
		b.WriteString("\n")
	}

	b.WriteString("Now summarize the portfolio’s strengths, potential risks, and suggest if any stock looks over- or under-valued based on the fundamentals.\n")

	return b.String()
}

func GenerateMarkdownTable(enriched []csvparser.EnrichedHolding) string {
	var b strings.Builder

	b.WriteString("| Symbol | Qty | Avg Price | ROCE | PE | Growth | D/E |\n")
	b.WriteString("|--------|-----|-----------|------|----|--------|-----|\n")

	for _, h := range enriched {
		b.WriteString(fmt.Sprintf(
			"| %s | %d | ₹%.2f | %s | %s | %s | %s |\n",
			h.Symbol,
			h.Quantity,
			h.AvgPrice,
			emptyDash(h.ROCE),
			emptyDash(h.PE),
			emptyDash(h.ProfitGrowth),
			emptyDash(h.DebtToEquity),
		))
	}

	return b.String()
}

func emptyDash(s string) string {
	if strings.TrimSpace(s) == "" {
		return "-"
	}

	return s
}
