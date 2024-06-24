/***
Working on the Hourly rate and payment for the week computation
***/

package main

import (
	"errors"
	"fmt"
)

// provide some error values
var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid number of hours worked")
)

func payDay(hoursWorked, hourlyRate int) int {
	report := func() {
		fmt.Printf("Hours Worked: %d\nHourly Rate: %d\n", hoursWorked, hourlyRate)
	}
	defer report()

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		overTimePay := overTime * hourlyRate
		regularPay := 40 * hourlyRate
		return regularPay + overTimePay
	}

	return hoursWorked * hourlyRate
}

func main() {
	pay := payDay(82, 76)
	fmt.Printf("AyeDay PayDay: %d\n", pay)
}
