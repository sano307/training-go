package main

import (
	"fmt"
)

func main() {
	pow2 := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}

	for i, value := range pow2 {
		fmt.Printf("2^%d = %d\n", i, value)
	}

	binaryShift := make([]int, 10)
	for i, _ := range binaryShift {
		binaryShift[i] = 1 << uint(i)
	}

	for _, value := range binaryShift {
		fmt.Printf("%d\n", value)
	}
}
