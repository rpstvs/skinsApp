package Scrapper

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/rpstvs/skinsApp/database"
)

func (cfg *Configure) writeToDb(data *SearchResult) {
	ctx := context.Background()
	for _, item := range data.Results {
		cfg.mu.Lock()
		id, err := cfg.db.GetItemIDbyName(ctx, item.HashName)

		if err != nil {
			fmt.Println("item not in db creating one")
			cfg.db.CreateItem(ctx, database.CreateItemParams{
				ID:         uuid.New(),
				Itemname:   item.HashName,
				Imageurl:   BuildImageURL(item.AssetDescription.IconURL),
				Daychange:  0.00,
				Weekchange: 0.00,
			})
		}

		_, err = cfg.db.AddPrice(ctx, database.AddPriceParams{
			Pricedate: ConvertDate(),
			ItemID:    id,
			Price:     PriceConverter(item.SalePriceText),
		})

		if err != nil {
			log.Printf("Item: %s already updated", item.HashName)
			cfg.mu.Unlock()
			continue
		}

		cfg.UpdateChange(ctx, id)
		cfg.mu.Unlock()
		log.Printf("Updating Item: %s - Daily Change %s.2 \n", item.HashName, item.SalePriceText)
	}
}

func (cfg *Configure) UpdateChange(ctx context.Context, id uuid.UUID) {

	priceHistory, _ := cfg.db.GetPricebyId(ctx, id)

	dailyChange, weeklyChange, _ := PriceChange(priceHistory)

	err := cfg.db.UpdatePriceChange(ctx, database.UpdatePriceChangeParams{
		Daychange:  dailyChange,
		Weekchange: weeklyChange,
		ID:         id,
	})

	if err != nil {
		log.Println(err)
	}

}
