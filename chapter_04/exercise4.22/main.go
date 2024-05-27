/***
Update the doubler func from exercise21 to use a type switch and expand its abilities to deal with more types
***/

package main

import (
	"errors"
	"fmt"
)

// create doubler func to take in an interface and return a string and error
func doubler(v interface{}) (string, error) {
	// create a type switch using the arg
	switch t := v.(type) {
	// for string and bool, we are matching on one type ony and do not need to do any extra safety checks
	case string:
		return t + t, nil

	case bool:
		if t {
			return "truetrue", nil
		}
		return "falsefalse", nil

	// for floats, we are matching on more than one type
	// perform type assertion to be able to make it work
	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}
		// if type assertion fails, we get a panic.
		// but we can rely on the logic that only float32 can work directly with the result of type assertion
		return fmt.Sprint(t.(float32) * 2), nil

	// match on all the int and uint types
	case int:
		return fmt.Sprint(t * 2), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:
		return fmt.Sprint(t * 2), nil
	case int32:
		return fmt.Sprint(t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case uint:
		return fmt.Sprint(t * 2), nil
	case uint8:
		return fmt.Sprint(t * 2), nil
	case uint16:
		return fmt.Sprint(t * 2), nil
	case uint32:
		return fmt.Sprint(t * 2), nil
	case uint64:
		return fmt.Sprint(t * 2), nil

	default:
		return "", errors.New("unsupported data type passed")
	}
}

func main() {
	res, _ := doubler(-5)
	fmt.Println("-5 :", res)
	res, _ = doubler(5)
	fmt.Println("5 :", res)
	res, _ = doubler("yum")
	fmt.Println("yum :", res)
	res, _ = doubler(true)
	fmt.Println("true:", res)
	res, _ = doubler(float32(3.14))
	fmt.Println("3.14:", res)
}
