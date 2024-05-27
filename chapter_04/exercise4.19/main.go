/***
Define some structs and custom types
Embed thoese types into a struct
***/

package main

import "fmt"

// create a custom type
type name string

// create a struct called location with two int fields
type location struct {
	x int
	y int
}

// create a size struct with two int fields
type size struct {
	width  int
	height int
}

// create a struct named dot. This embeds each of the preceding structs
type dot struct {
	name
	location
	size
}

// create a func that returns a slice of dots
func getDots() []dot {
	// use the var notation, all values will have the zero values
	var dot1 dot

	// initialize dot2 with zero values
	dot2 := dot{}

	// set the name, use the type's name as if it were a field
	dot2.name = "A"
	// for size and location, use the promoted fields
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	// when initializing embedded types, you cannot use promotion
	// for name, the result is the same but for location and size, more work has to be put into this
	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width:  5,
			height: 7,
		},
	}

	// for dot4, we use the type name path to set the data
	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.width = 43

	// return all dots in a slice
	return []dot{dot1, dot2, dot3, dot4}
}

func main() {
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v : %#v\n", i+1, dots[i])
	}
}
