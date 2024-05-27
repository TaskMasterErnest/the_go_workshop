/*
NOTE: the function here has a thing that baffled me for a while
			A function was called and had another functio eevaluate the input of the first function
			findLongest has string inputs, when these are called, the getPassedArgs also activates and counts the input
			to check if they are of the minimum length required.
			Simply fascinating.

			Like this: call getLongest -> getLongest calls getPassedArgs -> getPassedArgs executed -> getLongest executes.
*/

package main

import (
	"fmt"
	"os"
)

// fucntion that takes in an int and returns a string
func getPassedArgs(minArgs int) []string {
	// check if the correct number of arguments are being passed
	if len(os.Args) < minArgs {
		fmt.Printf("Usage: At least %v arguments are needed", minArgs)
		os.Exit(1)
	}

	// jumping over the first arg
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

// function to loop over a slice and find the longest word
// two words of the same length found, print the first one
func findLongest(args []string) string {
	var longest string

	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func main() {
	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was : ", longest)
	} else {
		fmt.Println("There was an error!")
		os.Exit(1)
	}
}
