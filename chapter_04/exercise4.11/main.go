/***
With the make function, we will create some slices and display their length and capacity.
***/

package main

import "fmt"

// create function that returns three(3) int slices
func genSlices() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)

	return s1, s2, s3
}

func main() {
	// call the genSlices function, capture the returned values and print each slice's length and capacity
	s1, s2, s3 := genSlices()

	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s1: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s1: len = %v cap = %v\n", len(s3), cap(s3))
}
