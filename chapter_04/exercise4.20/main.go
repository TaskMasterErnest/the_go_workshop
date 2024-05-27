/***
In this exercise, we will do some numeric type conversions and intentionally cause some data issues
***/

package main

import (
	"fmt"
	"math"
)

// func that returns a string
func convert() string {
	// define variable to perform implicit tyoe conversion from math.MaxInt8 to int8
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	// convert smaller int type to larger int type. Always a safe operation
	m := fmt.Sprintf("int8 = %v > int64 = %v\n", i8, int64(i8))

	// convert from int to one(1) value above int8's max size. this will cause an overflow into int8's min size
	m += fmt.Sprintf("int = %v > int8 = %v\n", i, int8(i))

	// convert int8 into a float64, this does nto cause an overflow and the data is unchanged
	m += fmt.Sprintf("int8 = %v > float64 = %v\n", i8, float64(i8))

	// convert a float into an int. all decimal data is lost but the whole number is kept up
	m += fmt.Sprintf("float64 = %v = int %v\n", f64, int(f64))

	return m
}

func main() {
	fmt.Print(convert())
}
