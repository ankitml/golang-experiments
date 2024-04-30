package main

import (
	"fmt"
	"reflect"
)

const a string = "Ankit"

func type_flexible_func(i interface{}) {
	fmt.Println(i)

	t := reflect.TypeOf(i)
	fmt.Println(t)

	switch i.(type) { // can use t.(type) only in switch case.. but not outside it
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("default")
	}
}

// func main() {
// 	i := 5
// 	type_flexible_func(reflect.TypeOf(i))

// 	b := [...]int{1, 2, 3, 4, 5}
// 	fmt.Println(b)

// 	c := [...]int{100, 3: 400, 500}
// 	fmt.Println(c)
// }
