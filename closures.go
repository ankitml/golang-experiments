package main

import "fmt"

func typeseq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closures() {
	f := typeseq()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	g := typeseq()
	fmt.Println(g())
}

func main() {
	var f func(int) int

	f = func(n int) int {
		if n < 2 {
			return n
		}
		return f(n-1) + f(n-2)
	}

	fmt.Println(f(100))
}
