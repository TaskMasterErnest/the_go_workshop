/***
Define a struct
Add some fields
Create some values using different methods
***/

package main

import "fmt"

// define struct type in the package scope
// provide a unique name at package-level scope
type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

// create a function that returns a slice of the new struct
func getUsers() []user {
	// initialize user with the key-valu notation. a popular way of initializing structs
	u1 := user{
		name:    "Ernest",
		age:     25,
		balance: 1100.00,
		member:  false,
	}

	// another initialization method, order does not matter and fieds can be left out
	u2 := user{
		age:  17,
		name: "Peter Parker",
	}

	// possible to initialize a struct with values only.
	// in this way, ALL the fields must be present, and the order must match how they are originally defined
	u3 := user{
		"Bob Vance",
		45,
		130000.00,
		true,
	}

	// the var notation can be used to create a struct wit all fields being zero values
	// we can set the values with the dot(.) notation
	var u4 user
	u4.name = "Scott"
	u4.age = 43
	u4.balance = 17.09
	u4.member = true

	// return all the users wrapped in a slice
	return []user{u1, u2, u3, u4}
}

// in the main function, we get the slice of user and we loop over them
func main() {
	// copy over the data
	users := getUsers()

	for i := 0; i < len(users); i++ {
		fmt.Printf("%v: %#v\n", i, users[i])
	}
}
