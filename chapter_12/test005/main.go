/***
Parsing the headers from the CSV file to use
***/

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	i := `firstname, lastname, age
Cecelia, Frank, 24
Jimmy, Fitzsimmons, 17
Bello, Crooks, 25`

	r := csv.NewReader(strings.NewReader(i))

	// set headers to be true
	header := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if !header {
			for idx, value := range record {
				switch idx {
				case 0:
					fmt.Println("First Name:", value)
				case 1:
					fmt.Println("Last Name:", value)
				case 2:
					fmt.Println("Age:", value)
				}
			}
		}
		header = false
	}
}
