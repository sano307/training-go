package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	a, b := 0, 1

	// A select blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if multiple are ready.
	for {
		select {
		case c <- a:
			a, b = b, a+b
		case <-quit:
			fmt.Println("quit")
			return
		// Use a default case to try a send or receive without blocking.
		default:
			fmt.Println("Default case")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
