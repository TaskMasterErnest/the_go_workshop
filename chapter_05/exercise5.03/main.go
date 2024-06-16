/***
Create a fizzbuzz function with return values
Modify it so that it only accepts an integer
return two values, the number and a string of either fizz, buzz or fizzbuzz
***/

package main

import "fmt"

func fizzBuzz(i int) (int, string) {
	// check if number is present or given
	switch {
	case i%15 == 0:
		return i, "fizzbuzz"
	case i%3 == 0:
		return i, "fizz"
	case i%5 == 0:
		return i, "buzz"
	}

	return i, ""
}

func main() {
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		n, s := fizzBuzz(i)
		fmt.Printf("Results: %d	%s\n", n, s)
	}
}
