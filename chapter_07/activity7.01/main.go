/***
Calculate annual pay for managare and developer
Print out the individual pay for the year for each type
Developer will have the following:
	- tracking number of hours worked
	- reviews
Aim of this is to demonstrate polymorphism.
***/

package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        string
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual    Employee
	Salary        float64
	CommisionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) Pay() (string, float64) {
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName, d.HourlyRate * d.HoursWorkedInYear
}

func (m Manager) Pay() (string, float64) {
	fullName := m.Individual.FirstName + " " + m.Individual.LastName
	return fullName, m.Salary + (m.Salary * m.CommisionRate)
}

func payDetails(p Payer) {
	fullName, yearlyPay := p.Pay()
	fmt.Printf("%s got paid $%.2f for the year\n", fullName, yearlyPay)
}

func convertReviewToInt(str string) (int, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}

func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}

func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := OverallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.Individual.FirstName+" "+d.Individual.LastName, averageRating)
	return nil
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["Work Quality"] = 5
	employeeReview["Teamwork"] = 2
	employeeReview["Communication"] = "Good"
	employeeReview["Problem Solving"] = 4
	employeeReview["Dependability"] = "Fair"

	d := Developer{Individual: Employee{ID: "d-137", FirstName: "Anthon", LastName: "Joggins"}, HourlyRate: 45.0, HoursWorkedInYear: 2080.0, Review: employeeReview}

	payDetails(d)
}
