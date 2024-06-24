/***
Implenting more functionalities in the banking application
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
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}
	return nil
}

func (d *directDeposit) validateLastName() error {
	if d.lastName == "" {
		return ErrInvalidLastName
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
		lastName:      "KlugHammer",
		firstName:     "Ernest",
		bankName:      "Chase",
		routingNumber: 100,
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
