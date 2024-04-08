// the bubble sort, the hard way
package main

import "fmt"

func swap(digits []int) {
	// perform swapping
	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < len(digits)-1; i++ {
			if digits[i] > digits[i+1] {
				digits[i], digits[i+1] = digits[i+1], digits[i]
				swapped = true
			}
		}
	}
}

func main() {
	nums := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}

	// perform bubble sort
	swap(nums)

	fmt.Println(nums)
}
