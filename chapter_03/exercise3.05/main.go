// we'll declare a string and initialize it with a multi-byte string value.
// We'll then loop over the string using range to give us each character, one at a time.
// We'll then print out the byte index and the character to the console

package main

import "fmt"

func main() {
	// declare a string with a multi-byte string
	logLevel := "デバッグ"

	// loop over with range and capture index and value
	for index, runeValue := range logLevel {
		fmt.Println(index, string(runeValue))
	}
}
