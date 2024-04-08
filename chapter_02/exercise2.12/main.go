package main

import "fmt"

func main() {
	shapes := []string{"square", "circle", "triangle"}

	for key, value := range shapes {
		fmt.Printf("This is %v and its index is %v\n", value, key)
	}
}
