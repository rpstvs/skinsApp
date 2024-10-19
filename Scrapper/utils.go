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
