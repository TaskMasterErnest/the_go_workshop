package main

import "fmt"

// create a function and initialize with 9 values
func message() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// extract the first value directly as an int; a slice using low:high; using only high
	m := fmt.Sprintln("First : ", s[0], s[0:1], s[:1])

	// the last element
	m += fmt.Sprintln("Last : ", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])

	// get the first 5 values and add them to the message
	m += fmt.Sprintln("First 5 : ", s[:5])

	// the last 4 values
	m += fmt.Sprintln("Last 4 : ", s[5:])

	// extract 5 values from the middle
	m += fmt.Sprintln("Middle 5 : ", s[2:7])

	return m

}

func main() {
	fmt.Println(message())
}
