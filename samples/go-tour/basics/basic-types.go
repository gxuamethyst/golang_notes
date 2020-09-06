package main

import (
	"fmt"
	"math/cmplx"
)

/*
1. Golang asic types

bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64
complex64 complex128

2. variables declared without an explicit initial value are given their zero value.
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/

var (
	// ToBe -- just for test
	ToBe   bool       // golint: exported var Xxx should have comment or be unexported
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("type: %T, value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T, value: %v\n", maxInt, maxInt)
	fmt.Printf("type: %T, value: %v\n", z, z)
}
