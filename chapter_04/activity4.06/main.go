package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	var i = true
// 	fmt.Printf("%v is %T\n", i, i)
// }

func typeChecker(v interface{}) (string, error) {
	// do a type switch for multiple value types
	switch t := v.(type) {
	case string:
		return fmt.Sprintf("%v is %T", t, t), nil

	case int, int32, int64:
		if i, ok := t.(int64); ok {
			return fmt.Sprintf("%v is %T", i, i), nil
		}

		if i, ok := t.(int32); ok {
			return fmt.Sprintf("%v is %T", i, i), nil
		}

		return fmt.Sprintf("%v is %T", t, t), nil

	case bool:
		return fmt.Sprintf("%v is %T", t, t), nil

	default:
		return "", errors.New("unsupported data type passed")

	}
}

func main() {
	res, _ := typeChecker(int32(8))
	fmt.Println(res)
	res, _ = typeChecker(true)
	fmt.Println(res)
	res, _ = typeChecker("blues")
	fmt.Println(res)
}
