package Scrapper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/rpstvs/skinsApp/database"
)

func (cfg *Configure) Get_skins(start int) {

	defer cfg.wg.Done()
	//time.Sleep(10 * time.Second)
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
		log.Printf("HTTP STATUS CODE: %d - Results is nil, skipping", resp.StatusCode)
		return
	}

	for _, item := range results.Results {
		x, err := cfg.db.CreateItem(context.Background(), database.CreateItemParams{
			ID:         uuid.New(),
			Itemname:   item.HashName,
			Imageurl:   BuildImageURL(item.AssetDescription.IconURL),
			Daychange:  0.00,
			Weekchange: 0.00,
		})
		log.Printf("Added Item: %s and err: %s \n", x.Itemname, err)
	}

}
