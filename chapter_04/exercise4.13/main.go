/***
Define a map with some data and add a new element to it.
Print the map to the console
***/

package main

import "fmt"

// create a function that returns a map with string keys and string values
func getUsers() map[string]string {
	// put in some elements
	users := map[string]string{
		"58":  "Ernest",
		"98":  "Taskmaster",
		"100": "TaskmasterErnest",
	}

	// add a new user user
	users["53"] = "Michaelangelo"

	return users
}

func main() {
	fmt.Println("Users: ", getUsers())
}
