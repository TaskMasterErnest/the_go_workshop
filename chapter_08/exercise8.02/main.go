/***
A function that will load stuff into a map before the main function runs
This makes use of init statements
***/

package main

import (
	"fmt"
)

var budgetCategories = make(map[int]string)

func init() {
	fmt.Println("Initializing Budget Categories")
	budgetCategories[1] = "Car Insurance"
	budgetCategories[2] = "Mortgage"
	budgetCategories[3] = "Electricity"
	budgetCategories[4] = "Retirement"
	budgetCategories[5] = "Vacation"
	budgetCategories[6] = "Groceries"
	budgetCategories[7] = "Car Payment"
}

func main() {
	for k, v := range budgetCategories {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}
}
