package main

import (
	"fmt"
)

// A var statement can be at package or function level.

var result bool
var str string

func main() {
	var i int
	var j int = 3
	// Inside a function,
	// the := short assignment statement can be used in place of a var declaration with implicit type.
	// Outside a function,
	// every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	k := 24
	var a, b, c = true, 1, "test"
	fmt.Println(result, str, i, j)
	fmt.Println(a, b, c)
}
