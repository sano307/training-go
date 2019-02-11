package main

import (
	"fmt"
	"math/cmplx"
)

const epsilon = 1e-9

func main() {
	fmt.Println(epsilon == 0.000000001) // true
	fmt.Println(cmplx.Pow(1+2i, 1.0/3))
	fmt.Println(cbrt(1 + 2i))
}

func cbrt(x complex128) complex128 {
	z := 1 + 0i
	p := z

	for {
		z = z - (z*z*z-x)/(3*z*z)

		if cmplx.Abs(z-p) < epsilon {
			break
		}

		p = z
	}

	return z
}
