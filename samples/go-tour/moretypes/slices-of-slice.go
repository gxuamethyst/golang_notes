package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"0", "0", "0"},
		[]string{"0", "0", "0"},
		[]string{"0", "0", "0"},
	}
	board[0][0] = "1"
	board[1][1] = "1"
	board[2][2] = "1"
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}
