package main

import "fmt"

func main() {
	// declare variables with an initial value
	query, limit, offset := "bat", 10, 0

	// change the variable values using one-line mgt
	query, limit, offset = "ball", offset, 20

	fmt.Println(query, limit, offset)
}
