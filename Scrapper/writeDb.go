package Scrapper

import (
	"context"

	"github.com/google/uuid"
	"github.com/rpstvs/skinsApp/database"
)

func (cfg *Configure) writeToDb(data *SearchResult) {

	ctx := context.Background()
	for _, item := range data.Results {
		cfg.db.CreateItem(ctx, database.CreateItemParams{
			ID:         uuid.New(),
			Itemname:   item.HashName,
			Imageurl:   BuildImageURL(item.AssetDescription.IconURL),
			Daychange:  0.00,
			Weekchange: 0.00,
		})

		id, _ := cfg.db.GetItemIDbyName(ctx, item.HashName)
		cfg.db.AddPrice(ctx, database.AddPriceParams{
			Pricedate: ConvertDate(),
			ItemID:    id,
			Price:     PriceConverter(item.SalePriceText),
		})

		priceHistory, _ := cfg.db.GetPricebyId(ctx, id)

		dailyChange, weeklyChange, monthlyChange := PriceChange(priceHistory)

		//log.Printf("Added Item: %s and err: %s \n", x.Itemname, err)
	}
}
