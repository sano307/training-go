package main

import (
	"fmt"
	"math"
)

// Possible types: Character, String, Boolean, or Number
const company = "Cyberagent"

var year, month, day int = 2019, 2, 10

func main() {
	fmt.Println("Circular constant: ", math.Pi)

	fmt.Println("1 + 2 =", add(1, 2))

	a, b := swap("Kim", "Inseo")
	fmt.Println(a, b)

	fmt.Println(split(29))

	fmt.Println(year, month, day)

	hour := 0
	fmt.Println(year, month, day, hour)

	fmt.Println("Company name:", company)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	score := 1
	for score < 1000 {
		score += score
	}
	fmt.Println(score)

	// Infinite loop
	// for {
	//     To do something...
	// }

	fmt.Println(squt(2), squt(-4))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 10))
}

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func squt(x float64) string {
	if x < 0 {
		return squt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x float64, n float64, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}
