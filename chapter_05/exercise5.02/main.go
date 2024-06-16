/***
Get a slice of columen headers from a CSV file
Map the index values to column headers
***/

package main

import (
	"fmt"
	"strings"
)

func csvHdrCol(header []string) {
	//initiate a map of index to the column header, int to strings
	csvHeadersToColumnIndex := make(map[int]string)

	// range over the header to process each string
	for i, v := range header {
		// remove trailing whitespaces from value
		v = strings.TrimSpace(v)

		// v = strings.ToLower(v)
		// csvHeadersToColumnIndex[i] = v

		// switch statement for matching; lower all cases for exact matching
		// get the header, set index value for header in map
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	// return the map, not necessary in real-world code but ...
	fmt.Println(csvHeadersToColumnIndex)
}

func main() {
	hdr := []string{"emp_id", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr)

	hdr2 := []string{"employee", "emp_id", "address", "hourly rate", "hours worked", "manager"}
	csvHdrCol(hdr2)
}
