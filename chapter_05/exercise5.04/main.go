/***
For the CSV exercise, we are now going to return the map as the result.
The map being returned is the index-to-column header mapping
***/

package main

import (
	"fmt"
	"strings"
)

func csvColHdr(hdr []string) map[int]string {
	csvIdxToCol := make(map[int]string)

	// process ieach stirng in the slice
	for i, v := range hdr {
		// remove trailing spaces in string
		v = strings.TrimSpace(v)

		// make the strings all lowercase
		switch strings.ToLower(v) {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return csvIdxToCol
}

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvColHdr(hdr)
	fmt.Println("Result: ", result)
}
