/***
After some initial alpha testing, the bank no longer wants the application to crash
We instead need to recover from the panic and print the error that caused the panic
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

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() error {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidRoutingNumber {
				fmt.Println(r)
			}
		}
	}()

	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
	return nil
}

func (d *directDeposit) validateLastName() error {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidLastName {
				fmt.Println(r)
			}
		}
	}()

	if d.lastName == "" {
		panic(ErrInvalidLastName)
	}
	return nil
}

func (d *directDeposit) report() {
	fmt.Println("****************************************************\n")
	fmt.Println("Last Name: ", d.lastName)
	fmt.Println("First Name: ", d.firstName)
	fmt.Println("Bank Name: ", d.bankName)
	fmt.Println("Routing Number: ", d.routingNumber)
	fmt.Println("Account Number: ", d.accountNumber)
}

func main() {
	me := directDeposit{
		lastName:      "",
		firstName:     "Ernest",
		bankName:      "Chase",
		routingNumber: 80,
		accountNumber: 12367495,
	}

	err := me.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	err = me.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}

	me.report()
}
