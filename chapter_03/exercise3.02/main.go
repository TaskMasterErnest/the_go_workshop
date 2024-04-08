package main

import "fmt"

func main() {
	// initialize an integer with int
	var a int = 100

	// initialize an integer with float32
	var b float32 = 100

	// initialize an integer with float64
	var c float64 = 100

	// divide each one by 3
	fmt.Println(a / 3)
	fmt.Println(b / 3)
	fmt.Println(c / 3)

	// get the numbers back by multiplication
	fmt.Println((a / 3) * 3)
	fmt.Println((b / 3) * 3)
	fmt.Println((c / 3) * 3)
}
