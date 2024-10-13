package main

import (
	"fmt"
	"time"
)

func main() {
	totalItems := getTotalItems()

	cfg := InitConfig(totalItems)

	for i := 0; i < totalItems; i += 100 {
		time.Sleep(1 * time.Second)
		fmt.Println("Starting a new thread %d", i)
		cfg.wg.Add(1)
		go cfg.get_skins(i)
	}

	go func() {
		cfg.wg.Wait()
		close(cfg.ch)
	}()

	allresults := make([]interface{}, 0)

	for _, results := range <-cfg.ch {
		allresults = append(allresults, results)
	}
	fmt.Println("Fetched %d items \n", len(allresults))
}
