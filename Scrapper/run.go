package Scrapper

import (
	"fmt"
	"time"
)

func (cfg *Configure) Run_Scrapper() {

	for i := 0; i < cfg.totalItems; i += 100 {
		time.Sleep(20 * time.Second)
		cfg.wg.Add(1)
		fmt.Println("Starting a new thread %d", i)

		go cfg.Get_skins(i)
	}

	go func() {
		cfg.wg.Wait()
		close(cfg.ch)
	}()

	allresults := make([]SearchResult, 0)
	fmt.Println("appending new result.")
	for result := range cfg.ch {
		fmt.Printf("Start: %d, TotalCount: %d\n", result.Start, result.TotalCount)
		allresults = append(allresults, result)
	}

	fmt.Println(len(allresults))
}
