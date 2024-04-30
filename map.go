package main

import (
	"fmt"
	"reflect"
)

func aboutMaps() {
	m := make(map[string]string)

	m["f"] = "fafda"
	m["d"] = "dhokla"

	fmt.Println("map m:", m)

	m1 := m["f"]
	fmt.Println("m1", m1, reflect.TypeOf(m1))

	m2 := m["g"]
	fmt.Println("m2", m2, reflect.TypeOf(m2))

	a, b := m["a"]
	fmt.Println("a, b", a, b) // b should have false and a should be empty string

	a, b = m["f"]
	fmt.Println("a, b", a, b) // b should have true and a should be "fafda"

	n := map[string]int{"fafda": 1, "dhokla": 2}
	fmt.Println(n)

}
