/***
Create a function that tells the difference in time between the current time zone and a specified time zone
Use LoadLocation() to specify the time zone
Use In() to convert the time
Output should be in ANSIC format
***/

package main

import (
	"fmt"
	"time"
)

func timeDiff(timezone string) (string, string) {
	current := time.Now()
	remoteZone, _ := time.LoadLocation(timezone)
	remoteTime := current.In(remoteZone)

	fmt.Println("The current time is:", current.Format(time.ANSIC))
	fmt.Println("The time in:", remoteZone, "is:", remoteTime)

	return current.Format(time.ANSIC), remoteTime.Format(time.ANSIC)
}

func main() {
	fmt.Println(timeDiff("America/New_York"))
}
