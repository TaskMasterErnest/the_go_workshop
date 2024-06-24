/***
Causing a runtime error and expecting a crash
***/

package main

import "fmt"

func main() {
	nums := []int{2, 4, 5, 9, 3}
	// wrong looping to sum all numbers in slice; presents Runtime Error
	// total := 0
	// for i := 0; i <= 10; i++ {
	// 	total += nums[i]
	// }
	// fmt.Println("Total: ", total)

	// right implemetaton
	total := 0
	for i := range nums {
		total += nums[i]
	}
	fmt.Println("Total: ", total)
}
