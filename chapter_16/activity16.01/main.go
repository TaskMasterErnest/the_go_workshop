package main

import (
	"fmt"
	"log"
	"sync"
)

func sumInBits(from int, to int, wg *sync.WaitGroup, mu *sync.Mutex, result *string) {
	defer wg.Done()
	for i := from; i <= to; i++ {
		mu.Lock()
		*result = fmt.Sprintf("%s|%v|", *result, i)
		mu.Unlock()
	}
	return
}

func main() {
	var result string
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go sumInBits(1, 25, wg, mu, &result)
	go sumInBits(26, 50, wg, mu, &result)
	go sumInBits(51, 75, wg, mu, &result)
	go sumInBits(76, 100, wg, mu, &result)
	wg.Wait()
	log.Println(result)
}
