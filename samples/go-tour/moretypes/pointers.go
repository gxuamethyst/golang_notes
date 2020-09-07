package main

import (
	"fmt"
)

func main() {
	var i int = 24
	var p *int = &i
	fmt.Println(i, *p, p)
	*p = 32
	fmt.Println(i, *p, p)
}
