/***
Create a function that will not have any parameters or return types.
The function wil iterate over a map and print the name ane number of items sold in the map
***/

package main

import "fmt"

// create itemsSold function
func itemsSold() {
	// initialize items map
	items := make(map[string]int)
	items["John"] = 42
	items["Celene"] = 109
	items["Micah"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println("is below expectations")
		} else if v > 40 && v <= 100 {
			fmt.Println("meets expectations")
		} else if v > 100 {
			fmt.Println("exceeded expectations")
		}
	}
}

func main() {
	// call function for itemsSold
	itemsSold()
}
