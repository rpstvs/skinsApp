package Scrapper

import (
	"log"
	"time"
)

func (cfg *Configure) Run_Scrapper() {

	for i := 0; i < cfg.totalItems; i += 100 {
		time.Sleep(15 * time.Second)
		cfg.wg.Add(1)
		log.Printf("Starting a new thread, starting index %d", i)

		go cfg.Get_skins(i)
	}

	cfg.wg.Wait()

	log.Println("Scrapping ended")

}
