package main

import (
	"fmt"
)

type point struct {
	x int
	y int
	z int
}

func main() {
	p := point{1, 2, 0}
	fmt.Println(p)
	p.x = 3
	fmt.Println(p)

	addr := &p
	addr.y = 4
	fmt.Println(p)

	p1 := point{3, 4, 5}
	p2 := point{x: 3}
	p3 := point{}
	p4 := &point{z: 5}
	fmt.Println(p1, p2, p3, p4, *p4)
}
