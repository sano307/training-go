package main

import (
	"fmt"
	"math"
)

const epsilon float64 = 1e-9

func main() {
	a, aCount := sqrt(2)
	fmt.Println(a, "count:", aCount)
	fmt.Println(math.Sqrt(2))

	b, bCount := sqrt(0)
	fmt.Println(b, "count:", bCount)

	c, cCount := sqrt(0.0000001)
	fmt.Println(c, "count:", cCount)

	d, dCount := sqrt(123456789123456789)
	fmt.Println(d, "count:", dCount)
}

func sqrt(x float64) (float64, int) {
	z := 1.0
	p := z
	count := 0

	for {
		z = z - (z*z-x)/(2*z)

		if math.Abs(z-p) < epsilon {
			break
		}

		count++
		p = z
	}

	return z, count
}

/*
func sqrt(x float64) (float64, int) {
	var z, before, after float64 = 1.0, 0.0, 0.0

	count := 0
	for {
		before = z
		after = z - (z*z-x)/(2*z)

		if nearlyEqual(before, after, epsilon) {
			break
		}

		count++
		z = after
	}

	return z, count
}

func nearlyEqual(before, after, epsilon float64) bool {
	var absBefore = math.Abs(before)
	var absAfter = math.Abs(after)
	var diff = math.Abs(before - after)

	if before == after {
		// Handles infinities
		return true
	} else if before == 0 || after == 0 || diff < math.SmallestNonzeroFloat64 {
		// before or after is zero or both are extremely close to it
		return diff < (epsilon * math.SmallestNonzeroFloat64)
	} else {
		// Use relative error
		return diff/math.Min((absBefore+absAfter), math.MaxFloat64) < epsilon
	}
}
*/
