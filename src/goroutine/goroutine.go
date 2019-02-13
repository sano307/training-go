package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.
func main() {
	go say("world")
	say("Hello")
}

func say(word string) {
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(word)
	}
}
