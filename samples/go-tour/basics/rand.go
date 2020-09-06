package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("random number: ", rand.Intn(24)) // will be a number fixed.
	rand.Seed(time.Now().Unix())
	fmt.Println("random number: ", rand.Intn(24)) // random actually
}
