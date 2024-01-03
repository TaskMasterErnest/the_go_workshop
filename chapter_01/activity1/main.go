package main

import "fmt"

var firstName string
var familyName string
var Age int
var peanutAllergy bool

func main() {
	firstName := "Magdalene"
	familyName := "Smith"
	Age := 45

	fmt.Printf("%s\n%s\n%d\n%v\n", firstName, familyName, Age, peanutAllergy)
}
