/***
Logic error with walking distances
Semantic error shown
***/

package main

import "fmt"

// incorrect logic
// func main() {
// 	km := 2
// 	if km > 2 {
// 		fmt.Println("Jump in the taxi.")
// 	} else {
// 		fmt.Println("Walkity walk walk")
// 	}
// }

// correct logic
func main() {
	km := 2
	if km >= 2 {
		fmt.Println("Jump in the taxi.")
	} else {
		fmt.Println("Walkity walk walk")
	}
}
