package main

import (
	"fmt"
	"math"
)

func main() {
	// Function Values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))

	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// Function Closures
func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}
