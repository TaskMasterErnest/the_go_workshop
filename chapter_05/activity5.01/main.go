/***
Calculate the working hours of employees for a week
Use that to calculate the payable salary amount
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

func main() {
	dev := Developer{}
	dev.LogHours(Monday, 8)
	dev.LogHours(Tuesday, 8)

	// print the total working hours
	fmt.Println("Hours worked on Monday: ", dev.WorkWeek[Monday])
	fmt.Println("Hours worked on Tuesday: ", dev.WorkWeek[Tuesday])
	fmt.Println("Hours worked on Wednesday: ", dev.WorkWeek[Wednesday])
	fmt.Printf("Hours worked this week: %d\n", dev.HoursWorked())
}
