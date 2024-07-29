/***

***/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("File Contents:")
	fmt.Println(string(content))

	r := strings.NewReader("No file here.")
	c, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Contents of strings.NewReader:")
	fmt.Println(string(c))
}
