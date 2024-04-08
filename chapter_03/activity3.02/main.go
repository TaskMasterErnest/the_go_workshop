// package main

// import "errors"

// type PersonDetails struct {
// 	CreditScore int
// 	Income      float64
// 	Loan_Amount float64
// 	Loan_Term   int
// }

// func loanCredibility(p *PersonDetails) (float64, bool, error) {
// 	// var interestRate float64
// 	// var monthlyPayment float64

// 	// check the loan term
// 	if p.Loan_Term%12 != 0 {
// 		return 0, false, errors.New("please choose a new loan term")
// 	}

// 	// check credit score
// 	if p.CreditScore >= 450 {
// 		interestRate := (15 / 100)
// 		monthlyPayment := (p.Income * (20 / 100))
// 		interestPayment := p.Loan_Amount * interestRate * float64(p.Loan_Term)
// 	}
// }

package main

import (
	"errors"
	"fmt"
)

// define constants for scores and ratios
const (
	goodScore      = 450
	lowScoreRatio  = 10
	goodScoreRatio = 20
)

// predefine the errors
var (
	ErrCreditScore = errors.New("invalud credit score")
	ErrIncome      = errors.New("income invalid")
	ErrLoanAmount  = errors.New("loan amount invalid")
	ErrLoanTerm    = errors.New("loan term not a multiple of 12")
)

// function to check loan details
func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm int) error {
	// base interest rate
	interest := 20.0

	if creditScore >= goodScore {
		interest = 15.0
	}

	// validate credit score
	if creditScore < 1 {
		return ErrCreditScore
	}

	// validate income
	if income < 1 {
		return ErrIncome
	}

	// validate loan amount
	if loanAmount < 1 {
		return ErrLoanAmount
	}

	// validate loan term
	if loanTerm < 1 || loanTerm%12 != 0 {
		return ErrLoanTerm
	}

	// convert interest rate
	rate := interest / 100

	// calculate the payable amount
	payment := ((loanAmount * rate) / float64(loanTerm)) + (loanAmount / float64(loanTerm))

	// total cost of loan
	totalInterest := (payment * float64(loanTerm)) - loanAmount

	// declare a variable for approval
	approved := false

	// add condition to check if income is more than the payment
	if income > payment {
		// percentage of income to be taken by loan payment
		ratio := (payment / income) * 100

		// for a good creditScore, allow for a higher ratio
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = true
		}
	}

	// print all data in the format
	fmt.Println("Credit Score: ", creditScore)
	fmt.Println("Income: ", income)
	fmt.Println("Loan Amount: ", loanAmount)
	fmt.Println("Loan Term: ", loanTerm)
	fmt.Println("Monthly Payment: ", payment)
	fmt.Println("Rate: ", interest)
	fmt.Println("Total Cost: ", totalInterest)
	fmt.Println("Approved: ", approved)
	fmt.Println("")

	return nil
}

func main() {
	// approving an applicant
	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 12)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
