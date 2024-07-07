/***
Printing decimal. binary and hexadecimal numbers from 1 to 255 and right-aligning the results
The decimal width should be set to 3, binary to 8 and hex to 2
The goal is to properly format the output of the data
***/

package main

import "fmt"

func main() {
	for i := 1; i <= 255; i++ {
		fmt.Printf("Decimal: %.3d Binary: %.8b Hex: %.2x\n", i, i, i)
	}
}
