package main

import "sync"

func InitConfig(total_items int) *Configure {

	ch := make(chan SearchResult, total_items/100)

	return &Configure{
		wg:         &sync.WaitGroup{},
		ch:         ch,
		totalItems: total_items,
	}
}
