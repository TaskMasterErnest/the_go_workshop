package payroll

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        string
	FirstName string
	LastName  string
}

func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := ConvertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}

func ConvertReviewToInt(str string) (int, error) {
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

type Payer interface {
	Pay() (string, float64)
}

func PayDetails(p Payer) {
	fullName, yearlyPay := p.Pay()
	fmt.Printf("%s got paid $%.2f for the year\n", fullName, yearlyPay)
}
