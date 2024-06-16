/***
Write a closure that decrements a given starting value
***/

package main

import "fmt"

func decrement(i int) func() int {
	return func() int {
		i--
		return i
	}
}

func main() {
	counter := 4
	// call decrement on counter
	f := decrement(counter)
	fmt.Println(f())
	fmt.Println(f())
}
