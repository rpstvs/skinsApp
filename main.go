package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rpstvs/skinsApp/Scrapper"
)

func main() {
	godotenv.Load(".env")
	total_items := Scrapper.GetTotalItems()
	cfg := Scrapper.InitConfig(total_items)
	cfg.Run_Scrapper()
	log.Println("Scrapper job finished. Database updated.")
}
