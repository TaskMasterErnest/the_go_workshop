// looping over arrays and slices
// In this exercise, we'll loop over a collection of strings.
// We'll be using a slice, but the loop logic will also be the same set of arrays.
// We'll define the collection; we'll then create a loop that uses the collection to control when to stop looping and a variable to keep track of where we are in the collection

package main

import (
	"fmt"
)

func main() {
	// define a slice of strings ang initialize it with data
	names := []string{"Bennie", "King", "Yoofi", "Mantey"}

	// in the statement, we use the lenght of the names collection to check if we are at the end collection
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
