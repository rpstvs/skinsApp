package Scrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetTotalItems() int {
	URL := fmt.Sprintf("https://steamcommunity.com/market/search/render/?query=&start=0&count=10&search_descriptions=0&norender=1&sort_column=popular&sort_dir=desc&appid=730")
	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	results := &SearchResult{}

	err = json.Unmarshal(data, &results)
	if err != nil {
		log.Println(err)
	}

	if results.TotalCount == 0 {
		GetTotalItems()
	}

	return results.TotalCount
}

func BuildImageURL(imageId string) string {
	tmp := "https://community.akamai.steamstatic.com/economy/image/"

	return tmp + imageId
}

func ConvertDate() time.Time {
	currentTime := time.Now()

	// Truncate to get DD-MM-YYYY
	currentDate := currentTime.Truncate(24 * time.Hour)

	return currentDate
}

func PriceConverter(priceStr string) float64 {
	if len(priceStr) > 7 {
		priceStr = strings.ReplaceAll(priceStr, "$", "")
		priceStr = strings.ReplaceAll(priceStr, ",", "")
		priceStr = strings.ReplaceAll(priceStr, "-", "0")

		price, err := strconv.ParseFloat(priceStr, 64)

		if err != nil {
			log.Println("error parsing price")
		}

		return price
	}
	priceStr = strings.ReplaceAll(priceStr, "$", "")
	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	priceStr = strings.ReplaceAll(priceStr, "-", "0")

	price, err := strconv.ParseFloat(priceStr, 64)

	if err != nil {
		log.Println("error parsing price")
	}

	return price
}

func PriceChange(prices []float64) (float64, float64, float64) {
	if len(prices) == 0 {
		return 0.00, 0.00, 0.00
	}

	dailyChange := ((prices[0] - prices[1]) / prices[1]) * 100

	WeeklyChange := func(x []float64) float64 {
		if len(x) < 7 {
			return 0.0
		}
		var sum float64
		for i := range x {
			sum += x[i]
		}
		average := sum / 7.0
		return ((prices[0] - average) / average) * 100
	}
	weeklychange := WeeklyChange(prices[:7])

	MonthlyChange := func(x []float64) float64 {
		if len(x) < 30 {
			return 0.0
		}
		var sum float64
		for i := range x {
			sum += x[i]
		}
		average := sum / 30
		return ((prices[0] - average) / average) * 100
	}
	monthlyChange := MonthlyChange(prices[:30])

	return dailyChange, weeklychange, monthlyChange
}
