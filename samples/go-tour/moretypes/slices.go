package main

import (
	"fmt"
)

func main() {
	a := [7]int{1, 3, 5, 7, 9, 11, 13}
	var s1 []int = a[1:3] // index --> [1, 3)
	s2 := a[2:5]
	s3 := a[1:]
	s4 := a[:3]
	fmt.Println(a, s1, s2, s3, s4)

	// Slices are like references to arrays
	s1[0] = 24
	s2[1] = 32
	fmt.Println(a, s1, s2)

	// reference: https://tour.golang.org/moretypes/9
	b := []bool{true, false, false} // --> 'xxx := [3]bool{true, false, false}'(array) and 'b := xxx[0:3]'(slice)
	fmt.Println(b)
}
