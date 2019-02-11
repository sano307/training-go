package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.ScaleCallByValue(5)
	fmt.Println(v)

	v.ScaleCallByReference(5)
	fmt.Println(v)
}

// Append method to "Vertex" struct
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Append method to "MyFloat" float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

// Passing by value
func (v Vertex) ScaleCallByValue(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Passing by pointer
func (v *Vertex) ScaleCallByReference(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
