package Scrapper

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/rpstvs/skinsApp/database"
)

func (cfg *Configure) writeToDb(data *SearchResult) {

	ctx := context.Background()
	for _, item := range data.Results {
		_, err := cfg.db.CreateItem(ctx, database.CreateItemParams{
			ID:         uuid.New(),
			Itemname:   item.HashName,
			Imageurl:   BuildImageURL(item.AssetDescription.IconURL),
			Daychange:  0.00,
			Weekchange: 0.00,
		})

		if err != nil {

			id, _ := cfg.db.GetItemIDbyName(ctx, item.HashName)

			cfg.db.AddPrice(ctx, database.AddPriceParams{
				Pricedate: time.Now().UTC(),
			})
		}
		log.Printf("Added Item: %s and err: %s \n", x.Itemname, err)
	}
}

//TODO TRUNCATE DATE
