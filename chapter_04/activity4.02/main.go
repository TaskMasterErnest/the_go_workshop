package main

import (
	"fmt"
	"os"
)

var Users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Gary",
	"073": "Tracy",
}

// check it the ID exists
func getID(id string) (string, bool) {
	var exists bool = true
	identity, exists := Users[id]
	return identity, exists
}

func main() {
	// check the arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: You need at least one argument")
		os.Exit(1)
	}

	// get the id number from the arguments
	identity, exists := getID(os.Args[1])

	if !exists {
		fmt.Printf("Error: the ID does not match any name in database\n")
		os.Exit(1)
	}

	fmt.Println("Hello!", identity)
}
