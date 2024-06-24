/***
A bank application wants to add custom errors when checking for last name and valid routing numbers
The direct deposit procedure allows for invalid names and routing numbers to be used
The bank wants a descriptive error message when these incidents occur
Create two descriptive error messages
***/

package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name provided")
	ErrInvalidRoutingNumber = errors.New("invalid routing number submitted")
)

func main() {
	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingNumber)
}
