/***
Reading from a map with direct access and a loop
Check if the key exists in the map
***/

package main

import (
	"fmt"
	"os"
)

// define a map and initialize it with data
func getUsers() map[string]string {
	return map[string]string{
		"1090": "Jake",
		"1738": "Fetty",
		"0550": "Bobby",
	}
}

func getUser(id string) (string, bool) {
	// get a copy of the Users map
	users := getUsers()

	// get the value from the Users map with the ID passed in
	// check if it exists too
	user, exists := users[id]

	return user, exists
}

func main() {
	// check if the argument has been passed in
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userID := os.Args[1]
	name, exists := getUser(userID)

	// if key is not found, print a message, and then print all the users using a range loop, exit afterwards
	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\n", userID)
		for key, value := range getUsers() {
			fmt.Println(" ID:", key, "Name: ", value)
		}
		os.Exit(1)
	}

	// if all is well, print the name of the user
	fmt.Println("Name: ", name)
}
