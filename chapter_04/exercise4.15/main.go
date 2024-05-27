/***
Define a map
Delete an element from it using user input
Print out the new 'smaller' map.
***/

package main

import (
	"fmt"
	"os"
)

// define a map in the package scope
var users = map[string]string{
	"1738": "Fetty",
	"1957": "Kwame",
	"0800": "Dunder",
	"0224": "Elorm",
}

// create a function to delete from the users map
// using the ID passed in as the key
func deleteUser(id string) {
	delete(users, id)
}

// grab the ID and pass it into the function
func main() {
	// get ID from args
	if len(os.Args) < 2 {
		fmt.Println("ID value not passed")
		os.Exit(1)
	}

	userID := os.Args[1]

	deleteUser(userID)

	fmt.Println("Users: ", users)

}
