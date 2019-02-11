package main

import (
	"fmt"
	"math"
)

const epsilon float64 = 1e-9

type ErrNegativeSqrt float64

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}

func sqrt(x float64) (float64, int, error) {
	if x < 0 {
		return 0, 0, ErrNegativeSqrt(x)
	}

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

	return z, count, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt nagative number: %f", e)
}
