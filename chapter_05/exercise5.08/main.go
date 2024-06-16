/***
An extensible solution to calculate the salary of developers and managers
***/

package main

import "fmt"

func salary(x, y int, f func(int, int) int) int {
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}

func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}

func main() {
	dev := salary(50, 2080, developerSalary)
	manager := salary(105000, 15750, managerSalary)

	fmt.Printf("Manager's pay: %d\n", manager)
	fmt.Printf("Developer's pay: %d\n", dev)
}
