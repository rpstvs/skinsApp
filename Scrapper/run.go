package Scrapper

import (
	"log"
	"time"
)

func (cfg *Configure) Run_Scrapper() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	numWorkers := 50
	jobs := make(chan int, 1000)

	//Spawning Go Routines for workerpool
	for i := 0; i < numWorkers; i++ {

		cfg.wg.Add(1)
		go cfg.Worker(i, jobs, ticker)
	}

	for i := 0; i < cfg.totalItems; i += 100 {
		jobs <- i
	}

	close(jobs)
	cfg.wg.Wait()

	log.Println("Scrapping finished.")

}

func (cfg *Configure) Worker(id int, jobs chan int, ticker *time.Ticker) {
	defer cfg.wg.Done()
	for i := range jobs {
		<-ticker.C
		log.Printf("Worker %d, requesting items starting %d", id, i)
		cfg.Get_skins(i)
	}
}
