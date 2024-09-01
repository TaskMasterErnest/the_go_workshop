/***
In this example, we will calculate the sum of all numbers betweeen 1 and 100 but with 4 Goroutines
***/

package main

import (
	"log"
	"sync"
	"sync/atomic"
)

func sum(from int, to int, wg *sync.WaitGroup, result *int32) {
	defer wg.Done()
	for i := from; i <= to; i++ {
		atomic.AddInt32(result, int32(i))
	}
}

func main() {
	var result int32
	wg := &sync.WaitGroup{}

	wg.Add(4)
	go sum(1, 25, wg, &result)
	go sum(26, 50, wg, &result)
	go sum(51, 75, wg, &result)
	go sum(76, 100, wg, &result)

	wg.Wait()
	log.Println(result)
}
