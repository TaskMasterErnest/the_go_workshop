/***
Implementing an interface
***/

package main

import "fmt"

// create a Speaker interface with a Speak method that returns a string
type Speaker interface {
	String() string
}

// create a person struct with name, age and isMarried as fields
type person struct {
	name      string
	age       int
	isMarried bool
}

// create a String() method for the person struct and return a string value.
// this will satisfy the Stringer{} interface
func (p person) String() string {
	return fmt.Sprintf("\n%v (%v years old).\nMarried?: %v", p.name, p.age, p.isMarried)
}

// create a Speak() method for the person type struct
func (p person) Speak() string {
	return "Hi, my name is " + p.name
}

func main() {
	p := person{
		name:      "Adelaide",
		age:       23,
		isMarried: false,
	}

	fmt.Println(p.Speak())
	fmt.Println(p)
}
