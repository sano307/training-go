package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	// An error occurs if you do not close the channel in range loop.
	// fatal error: all goroutines are asleep - deadlock!
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
	}

	// Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.
	close(c)
}
