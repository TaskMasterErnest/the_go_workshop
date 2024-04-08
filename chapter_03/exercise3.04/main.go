package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	// declare and initialize an int to the highest possible number it can take
	intA := math.MaxInt64

	intA = intA + 1

	// we not use the big package. A custom type that we will initialize with int's highest value too
	bigA := big.NewInt(math.MaxInt64)

	// adding 1 to it
	bigA.Add(bigA, big.NewInt(1))

	// print out results
	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int: ", intA)
	fmt.Println("Big Int: ", bigA.String())
}
