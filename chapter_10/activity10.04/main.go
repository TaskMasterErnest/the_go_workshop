/***
Calculating a future date and time
***/

package main

import (
	"fmt"
	"time"
)

func hour2Second(h int) int {
	return h * 3600
}

func minute2Second(m int) int {
	return m * 60
}

func second2Second(s int) int {
	return s
}

func main() {
	now := time.Now()
	fmt.Println("The current time is", now.Format(time.ANSIC))

	// time in 6 hours, 6 minutes and 6 seconds
	total_seconds := hour2Second(6) + minute2Second(6) + second2Second(6)
	elapsedTime := time.Duration(total_seconds) * time.Second

	fmt.Println("The time 6 hours, 6 minutes and 6 seconds from now will be:", now.Add(time.Duration(elapsedTime)).Format(time.ANSIC))
}
