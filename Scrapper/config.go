package Scrapper

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/rpstvs/skinsApp/database"
)

func InitConfig(total_items int) *Configure {
	dbUrl := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbUrl)
	queries := database.New(db)
	if err != nil {
		fmt.Println(err)
		log.Println("Error opening connection with DB")
	}
	ch := make(chan SearchResult, total_items/100)

	return &Configure{
		wg:         &sync.WaitGroup{},
		ch:         ch,
		totalItems: total_items,
		db:         queries,
		mu:         &sync.Mutex{},
	}
}
