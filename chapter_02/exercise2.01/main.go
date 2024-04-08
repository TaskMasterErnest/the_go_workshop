// simple exercise to using if statements to control logic that will run.
// this experiment tries to identify even or odd numbers using a modulus and if statements

package main

import (
	"fmt"
)

func main() {
	// define input
	input := 5

	// perform check
	if input%2 == 0 {
		fmt.Println(input, "is even")
	}

	if input%2 == 1 {
		fmt.Println(input, "is odd")
	}
}
