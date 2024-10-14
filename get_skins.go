package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (cfg *Configure) get_skins(start int) {

	defer cfg.wg.Done()
	URL := fmt.Sprintf("https://steamcommunity.com/market/search/render/?query=&start=%d&count=100&search_descriptions=0&norender=1&sort_column=popular&sort_dir=desc&appid=730", start)
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

	if results == nil {
		log.Println("Results is nil, skipping")
		return
	}

	cfg.ch <- *results
}
