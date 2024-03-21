package main

// create two functions
// one accepts a value, adds 5 to it and prints number to console
// the other accepts a number as a pointer, adds 5 to it and prints the result to a console

import (
	"fmt"
)

func add5Value(count int) {
	count += 5
	fmt.Println("add5Value  :", count)
}

func add5Pointer(count *int) {
	// dereference the pointer
	*count += 5
	fmt.Println("add5Pointer  :", *count)
}

func main() {
	var count int

	// the order of execution is important here.
	// swap the order and you realize different results that further illustrate
	// what the pointer does

	add5Value(count)
	fmt.Println("add5Value post :", count)

	add5Pointer(&count)
	fmt.Println("add5Pointer post :", count)
}
