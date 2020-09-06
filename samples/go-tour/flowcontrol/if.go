package main

import (
	"fmt"
	"math"
)

func n(i int) int {
	if i <= 1 {
		return 1
	}
	return i * n(i-1)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim) // v is reachable
	}
	// fmt.Println(v) // v is unreachable
	return lim
}

func main() {
	fmt.Println(n(3))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
