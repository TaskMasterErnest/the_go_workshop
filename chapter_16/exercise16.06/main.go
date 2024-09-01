/***
Sum numbers from everywhere
We have numbers from several sources and we want to add them all in one place
We will have 4 goroutines sending numbers in particular ranges, and the main routine will calculate the sum
***/

package main

import (
	"log"
	"time"
)

// send all numbers to the channel. sleep for a microsecond for another routine to pick up work
func push(from, to int, out chan int) {
	for i := from; i <= to; i++ {
		out <- i
		time.Sleep(time.Microsecond)
	}
}

func main() {
	s1 := 0
	ch := make(chan int, 100)

	go push(1, 25, ch)
	go push(26, 50, ch)
	go push(51, 75, ch)
	go push(76, 100, ch)

	// gather all numbers to add, create a loop of 100 cycles
	for c := 0; c < 100; c++ {
		// read number from the channel
		i := <-ch
		// log to see which numbers came first
		log.Println(i)
		// calculate sum and show results
		s1 += i
	}

	log.Println(s1)
}
