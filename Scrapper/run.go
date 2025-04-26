package Scrapper

import "log"

func (cfg *Configure) Run_Scrapper() {
	numWorkers := 50
	jobs := make(chan int, 1000)
	/*

		for i := 0; i < cfg.totalItems; i += 100 {
			time.Sleep(15 * time.Second)
			cfg.wg.Add(1)
			log.Printf("Starting a new thread, starting index %d", i)

			go cfg.Get_skins(i)
		}

		cfg.wg.Wait()

		log.Println("Scrapping ended")

	*/

	//Spawning Go Routines for workerpool
	for i := 0; i < numWorkers; i++ {
		cfg.wg.Add(1)
		go cfg.Worker(jobs)
	}

	for i := 0; i < cfg.totalItems; i += 100 {
		jobs <- i
	}

	close(jobs)
	cfg.wg.Wait()

	log.Println("Scrapping finished.")

}

func (cfg *Configure) Worker(jobs chan int) {
	defer cfg.wg.Done()
	for i := range jobs {
		cfg.Get_skins(i)
	}
}
