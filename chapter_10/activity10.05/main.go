/***
Printing the Local time in different time zones
***/

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	NY, _ := time.LoadLocation("America/New_York")
	LA, _ := time.LoadLocation("America/Los_Angeles")

	fmt.Println("My local current time is:", now.Format(time.ANSIC))
	fmt.Println("The time in New York is:", now.In(NY).Format(time.ANSIC))
	fmt.Println("The time in Los Angeles is:", now.In(LA).Format(time.ANSIC))
}
