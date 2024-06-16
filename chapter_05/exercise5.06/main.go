/***
Create an anonymous function that is going to have an argument passed to it
This function will be used to calculate the square root of a number
***/

package main

import "fmt"

func main() {
	j := 9
	x := func(i int) int {
		return i * i
	}
	fmt.Printf("The square of %d is %d\n", j, x(j))
}
