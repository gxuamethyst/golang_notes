package main

import (
	"fmt"
)

type vertex struct {
	Lat float64
	Lon float64
}

// init map with value
var location = map[string]vertex{
	"Google": vertex{
		37.42202,
		-122.08408,
	},
	"Google_": {
		37.42202,
		-122.08408,
	},
}

func main() {
	// init map with make
	m := make(map[int]string)
	m[0] = "zero"
	m[1] = "one"
	fmt.Println(m[1])

	// insert
	location["Bell labs"] = vertex{
		40.68433,
		-74.39967,
	}
	fmt.Println(location)
	// get
	google := location["Google"]
	fmt.Println(google)
	// delete
	delete(location, "Google_")
	fmt.Println(location)
	// check if key existed
	v, ok := location["Google"]
	fmt.Println(v, ok)
}
