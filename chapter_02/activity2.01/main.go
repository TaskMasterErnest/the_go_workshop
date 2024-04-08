// implement fizzbuzz

package main

import (
	"fmt"
)

// function to hold the outputs
func output(n int) {
	if n%15 == 0 {
		fmt.Println("FizBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(n)
	}
}

func main() {
	// keep loop number
	var number int

	// loop that does 100 iterations
	for number = 1; number < 101; number++ {
		output(number)
	}
}
