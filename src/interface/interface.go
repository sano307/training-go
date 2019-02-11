package main

import (
	"fmt"
	"math"
	"os"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

// No explicit declaration of intent, no "implements" keyword.
type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

func main() {
	var abser Abser

	f := MyFloat(-math.Sqrt2)
	abser = f
	fmt.Println(abser.Abs())

	v := Vertex{3, 4}
	abser = &v
	fmt.Println(abser.Abs())

	var w Writer
	w = os.Stdout
	fmt.Fprintf(w, "Hello, World\n")
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
