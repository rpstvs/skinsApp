package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	totalItems := getTotalItems()

	maxProcs := runtime.GOMAXPROCS(0) // Get current setting
	fmt.Println("Max threads:", maxProcs)

	cfg := InitConfig(totalItems)

	for i := 0; i < 2000; i += 100 {
		time.Sleep(10 * time.Second)

		fmt.Println("Starting a new thread %d", i)
		cfg.wg.Add(1)
		go cfg.get_skins(i)
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
