/***
Working with concurrency
We will get the sum of two number ranges, 1 - 10 anf 1 - 100
We will like these calculations to happen independently and see both results at the same time.
***/

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func sum(from int, to int) int {
// 	result := 0
// 	for i := from; i <= to; i++ {
// 		result += i
// 	}
// 	return result
// }

// func main() {
// 	//first
// 	// s1 := sum(1, 10)
// 	// s2 := sum(1, 100)
// 	// fmt.Println(s1, s2)
// 	//second, a very crude way to wait for the concurrent function to finish executing.
// 	now := time.Now()
// 	var s1, s2 int
// 	go func() {
// 		s1 = sum(1, 100)
// 	}()
// 	time.Sleep(time.Second)
// 	s2 = sum(1, 10)
// 	log.Println(s1, s2)
// 	fmt.Println("...elapsed:", time.Since(now))
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func sum(from int, to int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 0
	for i := from; i <= to; i++ {
		result += i
	}
	fmt.Printf("Sum from %d to %d = %d\n", from, to, result)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	now := time.Now()
	go sum(1, 70, &wg)
	go sum(1, 2000, &wg)
	wg.Wait()
	fmt.Println("...elapsed:", time.Since(now))
	fmt.Println("...main done...")
}
