package main

import "fmt"

// a function that defines an array with out words
func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}

	// call the words in a specific order and return them
	// Sprintln allows us to capture the formatted text before it is printed
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

func main() {
	fmt.Println(message())
}
