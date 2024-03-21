// code is erroring out.
// code should print true but it's printing false
package main

import (
	"fmt"
)

func main() {
	count := 0

	if count < 5 {
		count = 10
		count++
	}
	fmt.Println(count == 11)
}
