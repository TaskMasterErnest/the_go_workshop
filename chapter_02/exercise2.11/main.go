// In this exercise, we'll use continue and break in a loop to show you how you can take control of it.
// We're going to create a loop that keeps going forever.
// This means we have to stop it with break manually.
// We'll also randomly skip loops with continue.
// We'll do this skipping by generating a random number, and if that number is divisible by 3, we'll skip the rest of the loop

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// create an empty loop
	for {
		// generate a random integrer between 0 and 8
		r := rand.Intn(8)

		// if the random number is divisible by 3, print "Skip" and skip the rest of the loop
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		}

		// if random number is divisible by 2, print "Stop" and stop using the loop
		if r%2 == 0 {
			fmt.Println("Stop")
			break
		}

		// print the number if it is neither of the two
		fmt.Println(r)
	}
}
