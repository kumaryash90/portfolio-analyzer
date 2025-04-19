package screener

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type HoldingDetail struct {
	Symbol       string
	ROCE         string
	PE           string
	ProfitGrowth string
	DebtToEquity string
}

func ScrapeFundamentals(symbol string) (HoldingDetail, error) {
	url := fmt.Sprintf("https://www.screener.in/company/%s/consolidated", strings.ToUpper(symbol))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return HoldingDetail{}, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return HoldingDetail{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return HoldingDetail{}, fmt.Errorf("req failed, status: %d", res.StatusCode)
	}

	// fmt.Println(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return HoldingDetail{}, err
	}

	var detail HoldingDetail
	detail.Symbol = symbol

	var whitespaceRegex = regexp.MustCompile(`\s+`)

	doc.Find("ul#top-ratios li").Each(func(i int, s *goquery.Selection) {

		labelRaw := s.Find("span.name").Text()
		valueRaw := s.Find("span[class*='value']").Text()

		label := strings.TrimSpace(whitespaceRegex.ReplaceAllString(labelRaw, " "))
		value := strings.TrimSpace(whitespaceRegex.ReplaceAllString(valueRaw, " "))

		value = strings.ReplaceAll(value, " %", "%")
		value = strings.ReplaceAll(value, "Cr. ", "Cr.")
		value = strings.ReplaceAll(value, "₹ ", "₹")

		// fmt.Printf("Found: %s -> %s\n", label, value)

		switch label {
		case "ROCE":
			detail.ROCE = value
		case "Stock P/E":
			detail.PE = value
		case "Debt to equity":
			detail.DebtToEquity = value
		}
	})

	doc.Find("table:contains('Compounded Profit Growth') tr").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "3 Years:") {
			value := strings.TrimSpace(s.Find("td").Last().Text())
			detail.ProfitGrowth = value
		}
	})

	return detail, nil
}
