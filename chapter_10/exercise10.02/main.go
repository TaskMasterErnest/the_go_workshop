/***
Craft a function that allows you to calculate the duration of the execution between two time variables
and returns a string that tells you how long the execution took place.
***/

package main

import (
	"fmt"
	"strconv"
	"time"
)

func elapsedTime(start time.Time, end time.Time) string {
	elapsed := end.Sub(start)
	Hours := strconv.Itoa(int(elapsed.Hours()))
	Minutes := strconv.Itoa(int(elapsed.Minutes()))
	Seconds := strconv.Itoa(int(elapsed.Seconds()))

	return "The total execution time elapsed is: " + Hours + " hour(s), " + Minutes + " minute(s) and " + Seconds + " second(s)!"
}

func main() {
	start := time.Now()
	sum := 0
	for i := 0; i < 2500000000000; i++ {
		sum += i
	}
	end := time.Now()
	fmt.Println(elapsedTime(start, end))
}
