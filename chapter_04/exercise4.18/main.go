/***
Define a comparable struct and create a value with it
Defiine and create values with anonymous structs that have the same structure as the already named struct
Compare them and print results to the console
***/

package main

import "fmt"

// define comparable struct
type point struct {
	x int
	y int
}

// create a func that returns two booleans
func compare() (bool, bool) {
	// create first anonymous struct
	point1 := struct {
		x int
		y int
	}{
		10,
		10,
	}

	// second anonymous struct, initialize to zero(0) and then change the value after initialization
	point2 := struct {
		x int
		y int
	}{}

	point2.x = 10
	point2.y = 5

	// final struct uses the names struct created originally
	point3 := point{10, 10}

	// return and compare them
	return point1 == point2, point1 == point3
}

// compare and print results
func main() {
	a, b := compare()

	fmt.Println("point1 == point2: ", a)
	fmt.Println("point1 == point3: ", b)
}
