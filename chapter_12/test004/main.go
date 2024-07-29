/***
Testing out working with CSV format with Go
***/

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	in := `firstname, lastname, age
Cecelia, Frank, 24
Jimmy, Fitzsimmons, 23
Raymar, Royce, 22`

	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(record)
	}
}
