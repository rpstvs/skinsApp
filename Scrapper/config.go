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

	return &Configure{
		wg:         &sync.WaitGroup{},
		totalItems: total_items,
		db:         queries,
		mu:         &sync.Mutex{},
	}
}
