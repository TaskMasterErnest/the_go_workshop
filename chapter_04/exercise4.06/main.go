package main

import "fmt"

func message() string {
	var m string

	arr := [4]int{1, 2, 3, 4}

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v : %v\n", i, arr[i])
	}

	return m
}

func main() {
	fmt.Println(message())
}
