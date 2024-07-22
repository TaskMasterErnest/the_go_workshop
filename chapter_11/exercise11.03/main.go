/***
For this program, the old JSON data being used is not well documented.
The data types are not known and neither is the structure.
Write a program to analyze the unknown JSON data structure, and for each field,
print the data type and the JSON key-value pair
***/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// perform type assertion to capture the unknown type
// each of the switch case statements should print the data type, key and value.
func main() {
	jsonData := []byte(`
	{
		"id": 2,
		"lname": "Washington",
		"fname": "Bill",
		"isEnrolled": true,
		"grades": [100, 86, 98, 87],
		"class": {
			"coursename": "World History",
			"coursenum": 101,
			"coursehours": 3
		}
	}`)

	// perform JSON validation
	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
		os.Exit(1)
	}

	// specify an empty interface
	var v interface{}
	// unmarshal the data
	err := json.Unmarshal(jsonData, &v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// performing type assertion on the content of the JSON received
	data := v.(map[string]interface{})
	for k, v := range data {
		switch value := v.(type) {
		case string:
			fmt.Println("(string):", k, value)
		case float64:
			fmt.Println("(float64):", k, value)
		case int:
			fmt.Println("(int):", k, value)
		case bool:
			fmt.Println("(bool):", k, value)
		case []interface{}:
			fmt.Println("(slice):", k, value)
			for i, j := range value {
				fmt.Println("  ", i, j)
			}
		default:
			fmt.Println("(unknown):", k, value)
		}
	}
}
