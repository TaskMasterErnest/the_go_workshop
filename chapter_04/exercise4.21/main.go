/***
Perform some type assertions and ensure all the safety checks are in place to do so
***/

package main

import (
	"errors"
	"fmt"
)

// create a func that accepts an interface and returns a string and an error
func doubler(v interface{}) (string, error) {
	// check if arg is an int, if yes, multiply by 2  an return it
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}

	// check if it a string, if yes, concatenate against itself and return it
	if s, ok := v.(string); ok {
		return s + s, nil
	}

	// if no matches, return an error and close the functions
	return "", errors.New("unsupported type passed")
}

// call doubler func on a variety of data
func main() {
	res, _ := doubler(5)
	fmt.Println("5: ", res)

	res, _ = doubler("yum")
	fmt.Println("yum: ", res)

	_, err := doubler(true)
	fmt.Println("true: ", err)
}
