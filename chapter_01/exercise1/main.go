package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// seed the random number generator
	rand.NewSource(time.Now().UnixNano())
	// generate a random number between 0 and add 1 to get a number between 1 and 5
	r := rand.Intn(5) + 1
	// generate stars based on number given
	stars := strings.Repeat("*", r)
	// print the stars
	fmt.Println(stars)
}
