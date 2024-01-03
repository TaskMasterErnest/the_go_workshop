package main

import "fmt"

func main() {
	// start billing
	// main course
	var total float64 = 2 * 13
	fmt.Println("Sub: ", total)

	// add another 4 items
	total = total + (4 * 2.25)
	fmt.Println("Sub: ", total)

	// discount 5 currency
	total = total - 5
	fmt.Println("Sub: ", total)

	// calculate tip
	tip := total * 0.1
	fmt.Println("Tip: ", tip)

	// add tip tp total
	total = total + tip
	fmt.Println("Total: ", total)

	// bill is split between 2 people
	split := total / 2
	fmt.Println("Split: ", split)

	// reward the visitor
	visitCount := 24
	visitCount = visitCount + 1

	// divide visit count by 5 dollars
	remainder := visitCount % 5

	if remainder == 0 {
		fmt.Println("With this visit, you have earned a reward.")
	}

}
