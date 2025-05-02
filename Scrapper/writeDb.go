package Scrapper

import (
	"context"
	"log"

	"github.com/rpstvs/skinsApp/database"
)

func (cfg *Configure) writeToDb(data *SearchResult) {
	ctx := context.Background()
	for _, item := range data.Results {
		cfg.mu.Lock()
		err := cfg.db.CreateItem(ctx, database.CreateItemParams{
			Classid:    item.AssetDescription.Classid,
			Itemname:   item.AssetDescription.MarketHashName,
			Imageurl:   BuildImageURL(item.AssetDescription.IconURL),
			Daychange:  0.00,
			Weekchange: 0.00,
		})

		if err != nil {
			log.Printf("Item: %s, class id: %s , err: %s \n", item.AssetDescription.MarketHashName, item.AssetDescription.Classid, err)
		}

		_, err = cfg.db.AddPrice(ctx, database.AddPriceParams{
			Pricedate: ConvertDate(),
			ItemID:    item.AssetDescription.Classid,
			Price:     PriceConverter(item.SalePriceText),
		})

		if err != nil {
			log.Printf("Item: %s already updated with err %s", item.HashName, err)
			cfg.mu.Unlock()
			continue
		}

		cfg.UpdateChange(ctx, item.AssetDescription.Classid)

		cfg.mu.Unlock()

		log.Printf("Updating Item: %s - Daily Change %s.2 \n", item.HashName, item.SalePriceText)
	}
}

func (cfg *Configure) UpdateChange(ctx context.Context, id string) {

	priceHistory, _ := cfg.db.GetPricebyId(ctx, id)

	dailyChange, weeklyChange, _ := PriceChange(priceHistory)

	err := cfg.db.UpdatePriceChange(ctx, database.UpdatePriceChangeParams{
		Daychange:  dailyChange,
		Weekchange: weeklyChange,
		Classid:    id,
	})

	if err != nil {
		log.Println(err)
	}

}
