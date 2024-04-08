// using switch statements with multiple case expressions

package main

import (
	"fmt"
	"time"
)

func main() {
	// specify a day
	bornDay := time.Sunday

	switch bornDay {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekday")

	case time.Saturday, time.Sunday:
		fmt.Println("Born on the weekend")

	default:
		fmt.Println("Error: day not valid")
	}
}
