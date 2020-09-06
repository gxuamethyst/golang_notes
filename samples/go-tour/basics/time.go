package main

import (
	"fmt"
	"time"
)

func main() {
	// in go-tour, the meaning of "2009-11-10 23:00:00 UTC" -- Go was publicly announced in November 2009
	// reference: https://en.wikipedia.org/wiki/Go_(programming_language)
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
}
