package main

import (
	"fmt"
	"time"
)

func main() {
	// using the switch initial statement to define the variable.
	switch bornDay := time.Sunday; {
	case bornDay == time.Sunday || bornDay == time.Saturday:
		fmt.Println("Born on the weekend!")

	default:
		fmt.Println("Born on some other day!")
	}
}
