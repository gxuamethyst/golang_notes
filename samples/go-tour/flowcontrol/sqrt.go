package main

import (
	"fmt"
)

// reference: https://en.wikipedia.org/wiki/Newton%27s_method
func sqrt(x float64) {
	z := x / 2.0
	for i := 1; i <= 5; i++ {
		fmt.Println("loop: ", i)
		fmt.Println("z before: ", z)
		z -= (z*z - x) / (2 * z)
		fmt.Println("z after: ", z)
	}
	fmt.Println("z result: ", z)
}

func main() {
	sqrt(2)
}
