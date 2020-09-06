package main

import (
	"fmt"
	"math"
)

func main() {
	// In Go, a name is exported if it begins with a capital letter.
	// For example, Pizza is an exported name, as is Pi,
	// which is exported from the math package.
	// pizza and pi do not start with a capital letter, so they are not exported.
	fmt.Println(math.Pi) // right
	// fmt.Println(math.pi) // wrong
}
