/***
Calculating payable amount for employees based on working hours

***/

package main

import "fmt"

// create Employee type
type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

// create Developer struct
type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Monday    Weekday = 0
	Tuesday   Weekday = 1
	Wednesday Weekday = 2
	Thursday  Weekday = 3
	Friday    Weekday = 4
	Saturday  Weekday = 5
	Sunday    Weekday = 6
)

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	total_hours := 0
	for _, hours := range d.WorkWeek {
		total_hours += hours
	}
	return total_hours
}

func (d *Developer) PayDay() (int, bool) {
	// check of total work hours per week > 40
	if d.HoursWorked() > 40 {
		// pay normal rate for weekly pay
		overtime := d.HoursWorked() - 40
		overtimeHours := overtime * 2
		overtimePay := d.HourlyRate * overtimeHours
		regularPay := d.HourlyRate * 40
		return overtimePay + regularPay, true
	}
	return d.HoursWorked() * d.HourlyRate, false
}

func (d *Developer) PayDetails() {
	for i, v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours:	", v)
		case 1:
			fmt.Println("Monday hours:	", v)
		case 2:
			fmt.Println("Tuesday hours:	", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours:	", v)
		case 5:
			fmt.Println("Friday hours:	", v)
		case 6:
			fmt.Println("Saturday hours:	", v)
		}
	}

	fmt.Println("Hours worked this week: ")
}

func main() {
	d := Developer{Individual: Employee{ID: 1, FirstName: "Tony", LastName: "Stark"}, HourlyRate: 10}
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	d.LogHours(Wednesday, 10)
	d.LogHours(Thursday, 10)
	d.LogHours(Friday, 6)
	d.LogHours(Saturday, 8)
	d.PayDetails()
}
