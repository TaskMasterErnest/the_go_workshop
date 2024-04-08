// Program to Measure Password Complexity
// In this exercise, we write a program for the portal to display whether the password entered meets the character requirements. The character requirements are as follows:
// Have a lowercase letter
// Have an uppercase letter
// Have a number
// Have a symbol
// Be 8 or more characters long

package main

import (
	"fmt"
	"unicode"
)

// a function that takes a string arg and returns a bool
func passwordChecker(pw string) bool {
	// convert the string into a rune
	pwR := []rune(pw)

	// count the number of multi-byte characters using len
	if len(pwR) < 8 {
		return false
	}

	// define some bool variables
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	// loop over the multi-byte characters, one at a time
	for _, v := range pwR {
		// use the unicode package to check whether the char is uppercase
		if unicode.IsUpper(v) {
			hasUpper = true
		}

		if unicode.IsLower(v) {
			hasLower = true
		}

		if unicode.IsNumber(v) {
			hasNumber = true
		}

		// for the symbols, we also include punctuations and use the or operator
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}

	}

	// return the variables
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	// enter a password
	var password string
	fmt.Print("Enter a password: ")
	fmt.Scan(&password)

	if passwordChecker(password) != true {
		fmt.Println("You password does not meet the conditions of a safe password. Pick a stronger one")
	} else {
		fmt.Println("Strong password! Proceed.")
	}
}
