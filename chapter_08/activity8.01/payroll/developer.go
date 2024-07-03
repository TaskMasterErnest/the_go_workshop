package payroll

import "fmt"

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName, d.HourlyRate * d.HoursWorkedInYear
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
