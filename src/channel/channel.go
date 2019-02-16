package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	c1 := make(chan int)
	go sum(numbers[:3], c1)
	go sum(numbers[3:], c1)
	x, y := <-c1, <-c1

	// 9, 6, 15
	fmt.Println(x, y, x+y)

	// Buffering channel
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	// fatal error: all goroutines are asleep - deadlock!
	// c2 <- 3
	fmt.Println(<-c2) // 1
	fmt.Println(<-c2) // 2
	c2 <- 3
	c2 <- 4
	fmt.Println(<-c2) // 3
	fmt.Println(<-c2) // 4
}

func sum(numbers []int, c chan int) {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	c <- sum
}
