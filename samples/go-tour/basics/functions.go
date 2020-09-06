package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func swap(x, y string) (string, string) { // multiple return
	return y, x
}

func split(sum int) (x, y int) { // named return values
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(24))
}
