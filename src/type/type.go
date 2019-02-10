package main

import "fmt"

/*
	int	 : int32 or int64 (Depend on environment)
	int8 : -2^7 ~ 2^7-1
	int16: -2^15 ~ 2^15-1
	int32: -2^31 ~ 2^31-1
	int64: -2^63 ~ 2^63-1

	uint  : uint32 or uint64 (Depend on environment)
	uint8 : 0 ~ 2^8-1
	uint16: 0 ~ 2^16-1
	uint32: 0 ~ 2^32-1
	uint64: 0 ~ 2^64-1
*/

var (
	isMale        bool    = true
	name          string  = "inseo"
	floor         int     = 5
	age           uint    = 27
	leftEyeSight  float32 = 1.0
	rightEyeSight float32 = 1.2
)

var (
	c1 complex64 = 1 + 2i
	c2 complex64 = 3 + 4i
)

func main() {
	if isMale {
		fmt.Println("I'm male")
	} else {
		fmt.Println("I'm female")
	}

	fmt.Println("My name is ", name)
	fmt.Println("I'm ", age)
	fmt.Println("My left eyesight is ", leftEyeSight, "and right eyesight is ", rightEyeSight)

	fmt.Println(c1 + c2)
	fmt.Println(c1 - c2)

	realPart := real(c1)
	imaginaryPart := imag(c1)
	fmt.Println(realPart, imaginaryPart)

	complexNumber := complex(realPart, imaginaryPart)
	fmt.Println(complexNumber)
}
