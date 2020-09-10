package main

import (
	"golang.org/x/tour/pic"
)

func picDemo(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy, dy)
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8(i * j)
		}
	}
	return result
}

// reference: https://tour.golang.org/moretypes/18
func main() {
	pic.Show(picDemo)
}
