/***
We are going to be validating Social Security Numbers.
The program will be accepting SSNs without the dashes
We will log the application process so we can trace the whole program progress and process
We do not want the application to stop when an invalid SSN is provided, it has to log the invalid number and
continue to the next
***/

package validate

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

var (
	ErrInvalidSSNLength  = errors.New("invalid SSN length, must be 9 characters exactly")
	ErrInvalidSSNNumbers = errors.New("invalid SSN character present")
	ErrInvalidSSNPrefix  = errors.New("invalid SSN prefix")
	ErrInvalidDigitPlace = errors.New("invalid SSN digit placement")
)

// checking length of provided string
func CheckLength(ssn string) error {
	if len(ssn) != 9 {
		log.Printf("the ssn: %s. Error: %s", ssn, ErrInvalidSSNLength)
	}
	return nil
}

// checking content of provided string, if it is numerals only
func numsOnly(str string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	if re.MatchString(str) {
		return true
	} else {
		return false
	}
}
func CheckNumbers(ssn string) error {
	if !numsOnly(ssn) {
		log.Printf("the ssn: %s. Error: %s", ssn, ErrInvalidSSNNumbers)
	}
	return nil
}

// checking the prefix of the provided string
func validPrefix(str string) bool {
	_, found := strings.CutPrefix(str, "000")
	return found
}
func CheckPrefix(ssn string) error {
	if validPrefix(ssn) {
		log.Printf("the ssn: %s. Error: %s", ssn, ErrInvalidSSNPrefix)
	}
	return nil
}

// checking the placement of digits
func validDigitPlace(str string) bool {
	_, found := strings.CutPrefix(str, "9")
	return found
}
func CheckDigitPlacement(ssn string) error {
	if validDigitPlace(ssn) {
		if string(ssn[3]) != "7" && string(ssn[3]) != "9" {
			log.Printf("the ssn: %s. Error: %s", ssn, ErrInvalidDigitPlace)
		}
	}
	return nil
}
