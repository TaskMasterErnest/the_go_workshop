/***
Enhance the payDay() function to recover from a panic
***/

package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid number of hours worked")
)

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Printf("hourly rate: %d\nerr: %v\n\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("hours workedL %d\nerr: %v\n\n", hoursWorked, r)
			}
		}
		fmt.Printf("Pay is calculated based on:\nHours Worked: %d\nHourly Rate: $%d\n", hoursWorked, hourlyRate)
	}()

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
	pay := payDay(40, 30)
	fmt.Printf("AyeDay PayDay: $%d\n", pay)
}
