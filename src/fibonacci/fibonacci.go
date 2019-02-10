package main

import (
	"fmt"
)

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}
}
