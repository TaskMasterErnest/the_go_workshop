/***
In this exercise, we are going to do the opposite of what was done in exercise 11.01.
We will marshal a struct into a JSON
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
	IsMarried     bool     `json:"-"`
	IsEnrolled    bool     `json:"enrolled,omitempty"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenumber"`
	Hours  int    `json:"coursehours"`
}

func newStudent(studentID int, lastName, middleInitial, firstName string, isMarried, isEnrolled bool) student {
	s := student{
		StudentID:     studentID,
		LastName:      lastName,
		MiddleInitial: middleInitial,
		FirstName:     firstName,
		IsMarried:     isMarried,
		IsEnrolled:    isEnrolled,
	}
	return s
}

func main() {
	s := newStudent(001, "Snorkel", "Amy", "Emily", false, false)
	student1, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student1))

	s2 := newStudent(002, "Sanchez", "Sheldon", "Robert", true, true)
	c := course{Name: "World Science", Number: 112, Hours: 4}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Introduction to Biology", Number: 328, Hours: 3}
	s2.Courses = append(s2.Courses, c)

	student2, err := json.MarshalIndent(s2, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student2))
}
