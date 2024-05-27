package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	// loop through args and append values to slice
	var args []string

	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

// function that accepts a slice of strings as a parameter and returns a slice of strings
func getLocals(extraLocals []string) []string {
	var locales []string

	locales = append(locales, "en_US", "fr_FR")

	locales = append(locales, extraLocals...)

	return locales
}

func main() {
	locales := getLocals(getPassedArgs())
	fmt.Println("Locales to use : ", locales)
}
