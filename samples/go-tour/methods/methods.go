package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x float64
	y float64
}

// A method is a function with a special receiver argument.
func (v vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.Abs())
}
