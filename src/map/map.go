package main

import (
	"fmt"
)

type Location struct {
	latitude  float32
	longitude float32
}

var locations map[string]Location

var cyberagent = map[string]Location{
	"AMoAd": Location{
		35.65816, 139.69862,
	},
	"AbemaTV": Location{
		35.65654, 139.69504,
	},
}

var parks = map[string]Location{
	"Yoyogi Park":   {35.67280, 139.69259},
	"Shinjuku Park": {35.68670, 139.70816},
}

func main() {
	locations = make(map[string]Location)

	locations["Google Japan"] = Location{
		35.66046, 139.72924,
	}
	fmt.Println(locations["Google Japan"])

	// var locations map[string]Location
	// panic: assignment to entry in nil map
	// locations["Google Japan"] = Location{
	//     35.66046, 139.72924,
	// }

	fmt.Println(cyberagent)
	fmt.Println(parks)

	basket := make(map[string]int)
	basket["Coke"] = 110
	fmt.Println(basket["Coke"])
	basket["Coke"] = 100
	fmt.Println(basket["Coke"])
	delete(basket, "Coke")
	fmt.Println(basket["Coke"])
	valueCoke, isCoke := basket["Coke"]
	fmt.Println(valueCoke, isCoke)

	basket["Water"] = 80
	valueWater, isWater := basket["Water"]
	fmt.Println(valueWater, isWater)
}
