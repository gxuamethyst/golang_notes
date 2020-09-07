package main

import (
	"fmt"
)

func main() {
	var a1 [10]int
	a1[0] = 11
	a1[2] = 24
	fmt.Println(a1)

	var a2 = [3]int{1, 3, 9}
	fmt.Println(a2)
}
