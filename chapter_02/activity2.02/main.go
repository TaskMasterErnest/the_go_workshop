// You have to find the word with the maximum count and print the word and its count using the following data

package main

import "fmt"

func main() {
	// map the words to the count
	lyrics := map[string]int{
		"Gonna": 3,
		"You":   5,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	var topWord string
	var topCount int

	// create a loop and use range to capture the word with the most count
	for key, value := range lyrics {
		if value > topCount {
			topCount = value
			topWord = key
		}
	}

	fmt.Println("Most popular word: ", topWord)
	fmt.Print("With a count of: ", topCount, "\n")
}
