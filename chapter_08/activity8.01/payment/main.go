package main

import (
	"fmt"
	"os"

	pay "github.com/the_go_workshop/chapter_08/activity8.01/payroll"
)

var employeeReview = make(map[string]any)

func init() {
	fmt.Println("Welcome to the Employee Pay and Performance Review")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func init() {
	employeeReview["Work Quality"] = 5
	employeeReview["Teamwork"] = 2
	employeeReview["Communication"] = "Good"
	employeeReview["Problem Solving"] = 4
	employeeReview["Dependability"] = "Fair"
}

func main() {
	d := pay.Developer{Individual: pay.Employee{ID: "d-137", FirstName: "Anthon", LastName: "Joggins"}, HourlyRate: 45.0, HoursWorkedInYear: 2080.0, Review: employeeReview}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pay.PayDetails(d)
}
