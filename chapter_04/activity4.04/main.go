package main

import "fmt"

// initialize slice with days of the week
// daysOfWeek := []string{
// 	"Monday",
// 	"Tuesday",
// 	"Wednesday",
// 	"Thursday",
// 	"Friday",
// 	"Saturday",
// 	"Sunday"
// }

func main() {
	daysOfWeek := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	fmt.Println(daysOfWeek)

	// change the order of the slice so that Sunday starts, and ends with Saturday
	newOrder := daysOfWeek[:len(daysOfWeek)-1]
	newOrder = append(daysOfWeek[len(daysOfWeek)-1:], newOrder...)
	fmt.Println(newOrder)
}
