// same example from exercise2.03 but with a twist.
// more rules about what numbers to check are added.
// move all validation logic to a function that returns an error; check if the error is nil , if not deal with the error

package main

import (
	"errors"
	"fmt"
)

// function to perform the validation; takes a single integer and returns an error
func validate(input int) error {
	// define some rules; if any are true,  return error using the New function in error package
	if input < 0 {
		return errors.New("Input cannot be a negative number")
	} else if input > 100 {
		return errors.New("Input cannot be greater than 100")
	} else if input%7 == 0 {
		return errors.New("Input cannot be divisible by 7")
	} else {
		return nil
	}
}

func main() {
	input := 20

	// call the function using the initial statement
	// use short variable assignment to capture the returned error
	// in the boolean expression, check that the error is not equal to nil
	if err := validate(input); err != nil {
		fmt.Println(err)
	} else if input%2 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}
