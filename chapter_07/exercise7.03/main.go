/***
Analyzing empty interface{} data
We a given a map, its key is a string and the value is an empty interface{}
The map's value contains different types of data stored in the value portion of the map
Determine each key's value type
This program will analyze the map[string]interface{}
Store the information about these types in a slice of structs that will hold the key name, data and the data type
***/

package main

import "fmt"

type record struct {
	key       string
	valueType string
	data      interface{} // interface{} because we want to store various data types that we do not know of
}

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

// newRecord() function, takes in key and interface and returns a record type
func newRecord(key string, i interface{}) record {
	// initialize a record{} type and assign it to a variable
	r := record{}
	r.key = key

	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person:
		r.valueType = "person"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}

func main() {
	// initialize the map
	m := make(map[string]interface{})
	a := animal{name: "adelaide", category: "feline"}
	p := person{lastName: "arthur", age: 32, isMarried: false}

	m["person"] = p
	m["animal"] = a
	m["age"] = 54
	m["isMarried"] = true
	m["lastName"] = "smith"

	// initialize a slice of records. Iterate over the map and add records to rs
	rs := []record{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}

	for _, v := range rs {
		fmt.Println("Key: ", v.key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}
}
