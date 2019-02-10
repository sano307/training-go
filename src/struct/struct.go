package main

import (
	"fmt"
)

type EyeSight struct {
	Left  float32
	Right float32
}

var (
	es1 = EyeSight{2.0, 2.0}
	es2 = &EyeSight{2.0, 2.0}
	es3 = EyeSight{Left: 0.5}
	es4 = EyeSight{}
)

func main() {
	fmt.Println(EyeSight{1.0, 1.2})

	es := EyeSight{1.0, 1.2}
	es.Left = 1.2
	fmt.Println(es.Left, es.Right)

	copy := &es
	copy.Right = 1.5
	fmt.Println("original:", es)
	fmt.Println("copy:", copy)

	fmt.Println(es1)
	fmt.Println(es2)
	fmt.Println(es3)
	fmt.Println(es4)
}
