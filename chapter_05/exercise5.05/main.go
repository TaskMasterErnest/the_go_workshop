/***
Sum up a variable number of arguments
Pass the arguments as a list or arguments and a slice
***/

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	i := []int{1, 5, 6, 7}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}
