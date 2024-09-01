package main

import (
	"fmt"
	"sync"
	"time"
)

func sum(from int, to int) int {
	result := 0
	for i := from; i <= to; i++ {
		result += i
	}
	return result
}

func main() {
	now := time.Now()
	var wg sync.WaitGroup
	var s1 int
	wg.Add(1)

	go func() {
		defer wg.Done()
		s1 = sum(1, 100)
	}()

	s2 := sum(1, 10)

	wg.Wait()
	fmt.Println(s1, s2)

	fmt.Println("...elapsed:", time.Since(now))
}
