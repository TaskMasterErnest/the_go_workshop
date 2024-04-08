package main

import "fmt"

func main() {
	// declare the int8 and uint8 values as close as possible
	var a int8 = 125
	var b uint8 = 253

	// loop through 5 times
	for i := 0; i < 5; i++ {
		a++
		b++

		// print the values
		fmt.Printf("%v) int8: %v\tuint8: %v\n", i, a, b)
	}
}
