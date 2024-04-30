package main

import "fmt"

func swap(a *int, b *int) {
	*b, *a = *a, *b
}

func swapper() {
	i := 1
	j := 2
	fmt.Println("i:", i, "j:", j)
	swap(&i, &j)
	fmt.Println("i:", i, "j:", j) // should print "i: 2 j: 1"
}
