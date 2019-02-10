package main

import (
	"fmt"
)

func main() {
	primeNumbers := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}
	fmt.Println(primeNumbers)

	count := len(primeNumbers)
	for i := 0; i < count; i++ {
		fmt.Printf("primeNumbers[%d] == %d\n", i, primeNumbers[i])
	}

	fmt.Println(primeNumbers[0:0])
	fmt.Println(primeNumbers[0:1])
	fmt.Println(primeNumbers[2:5])
	fmt.Println(primeNumbers[2:])

	dice := make([]int, 6)
	fmt.Println(dice)
	// Length of a slice.
	fmt.Println(len(dice))

	// The capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice.
	fmt.Println(cap(dice))

	var empty []int
	fmt.Println(empty, len(empty), cap(empty))
	fmt.Println(empty == nil)
}
