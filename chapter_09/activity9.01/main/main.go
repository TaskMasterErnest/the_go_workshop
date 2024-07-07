/***
COMMENT: Modify this code to be able to read from files; CSV format, JSON format or whatever.
***/

package main

import (
	"log"

	"github.com/the_go_workshop/chapter_09/activity9.01/validate"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	validateSSN := []string{"1234A4568", "1209834571", "113450961", "91260913", "90091234"}

	for index, ssn := range validateSSN {
		log.Printf("Validating data %d of %v", index+1, ssn)

		err := validate.CheckLength(ssn)
		if err != nil {
			log.Print(err)
		}

		err = validate.CheckNumbers(ssn)
		if err != nil {
			log.Print(err)
		}

		err = validate.CheckPrefix(ssn)
		if err != nil {
			log.Print(err)
		}

		err = validate.CheckDigitPlacement(ssn)
		if err != nil {
			log.Print(err)
		}
	}
}
