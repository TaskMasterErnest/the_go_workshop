/***
In this activity, we're going to set up a slice with some data and with one element to remove.
Then, you need to work out how to do this.
There are many ways to get this done, but can you work out the most compact way?
***/

package main

import "fmt"

var words = []string{
	"Good",
	"Good",
	"Bad",
	"Good",
	"Good",
}

// check if the word is in the slice
func checkWord(s []string, w string) bool {
	for _, v := range s {
		if v == w {
			return true
		}
	}
	return false
}

func main() {
	mySlice := words
	wordToRemove := "Bad"

	fmt.Println(checkWord(mySlice, wordToRemove))

	if checkWord(mySlice, wordToRemove) {
		// create a new empty slice
		var newSlice []string
		for _, v := range mySlice {
			// if the value is not the word to remove, append to the new slice
			if v != wordToRemove {
				newSlice = append(newSlice, v)
			}
		}
		fmt.Println(newSlice)
	} else {
		fmt.Printf("Value not Found\nHere's the list: %v\n", mySlice)
	}
}
