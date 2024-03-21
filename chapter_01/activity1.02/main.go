package main

import (
	"fmt"
)

// your job is to finish some code a co-worker started. Here, we have some unfinished code for you to complete.
// Your task is to fill in the missing code, where the comments are to swap the values of a and b.
// The swap function only accepts pointers and doesn't return anything

func main() {
	a, b := 5, 10

	// call swap here
	swap(&a, &b)

	fmt.Println(a == 10, b == 5)
}

func swap(a *int, b *int) {
	// swap the values here
	*a, *b = *b, *a
}
