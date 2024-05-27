package main

import "fmt"

// a function that defines an array
func defineArray() [10]int {
	var arr [10]int
	return arr
}

func main() {
	// print the result, user %#v to get additional details about the result, like its type.
	fmt.Printf("%#v\n", defineArray())
}
