/***
A program that takes a JSON from a web request for college class enrollment
The program will marshal the JSON data into a Go struct
The JSON will contain data about a student and the courses they are taking
After unmarshalling the JSON, print out the struct for verification purposes
***/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentID     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"midinitial"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenumber"`
	Hours  int    `json:"coursehours"`
}

func main() {
	data := []byte(`
	{
		"id": 123,
		"lname": "Smith",
		"midinitial": null,
		"fname": "Morty",
		"enrolled": true,
		"classes": [
			{
			"coursename": "Intro to Mechanics",
			"coursenumber": 123,
			"coursehours": 4
			},
			{
				"coursename": "Intro to Materials",
				"coursenumber": 234,
				"coursehours": 3
			},
			{
				"coursename": "Applied Electricity",
				"coursenumber": 145,
				"coursehours": 3
			}
		]
	}`)

	if !json.Valid(data) {
		fmt.Printf("Invalid JSON data: %s", data)
		os.Exit(1)
	}

	var s student

	err := json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
}
