/*
**
We will perform a sum of numbers in a predefined number of routines for them to gather the result at the end.
Essentially, we want to create a function to add numbers and receive numbers from a channel at the end.
The function will not know the numbers it will be working with beforehand (there will be no fixed range)
Also the addition should not happen in the main function
**
*/

package main

import "log"

// function to perform partial addition
// a fixed set of routines will be running this, waiting for numbers
func worker(in chan int, out chan int) {
	var sum int
	for i := range in {
		sum += i
		out <- sum
	}
}

// the sum function.
// this has the number of workers and the usual range for the numbers to add
func sum(workers int, from int, to int) int {
	// run the requested number of workers
	out := make(chan int, workers)
	in := make(chan int, 4)
	// this creates two in/out channels and runs the number of workers set by the workers parameter
	for i := 0; i < workers; i++ {
		go worker(in, out)
	}

	// create loop to send out all numbers to the in channel
	for i := from; i <= to; i++ {
		in <- i
	}
	close(in)

	// the in channel spreads the numbers to the worker routines to handle the partial sum
	// we then receive the partial sums
	var sum int
	for i := 0; i < workers; i++ {
		sum += <-out
	}
	close(out)

	return sum
}

func main() {
	result := sum(100, 1, 100)
	log.Println(result)
}
