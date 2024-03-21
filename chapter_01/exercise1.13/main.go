package main

import (
	"fmt"
	"time"
)

func main() {
	// declare a pointer using the var statement
	var count1 *int

	// create a variable using new
	count2 := new(int)

	countTemp := 5

	count3 := &countTemp

	// create a pointer from some types without a temporary variable, using time struct
	t := &time.Time{}

	// print each out
	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)
}
