package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getTotalItems() int {
	URL := fmt.Sprintf("https://steamcommunity.com/market/search/render/?query=&start=0&count=10&search_descriptions=0&norender=1&sort_column=popular&sort_dir=desc&appid=730")
	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		fmt.Errorf("Error occurred: %s", err)
		os.Exit(-1)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Errorf("Error occurred: %s", err)
	}

	data, _ := io.ReadAll(resp.Body)

	defer resp.Body.Close()

	results := &SearchResult{}

	err = json.Unmarshal(data, &results)
	return results.TotalCount
}
