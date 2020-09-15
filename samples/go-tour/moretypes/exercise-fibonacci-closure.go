package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a1 := -1
	a2 := 1
	return func() int {
		n := a1 + a2
		a1 = a2
		a2 = n
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
