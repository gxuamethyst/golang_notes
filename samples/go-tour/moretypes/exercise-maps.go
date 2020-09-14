package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for _, value := range words {
		result[value]++
	}
	return result
}

func main() {
	wc.Test(wordCount)
}
