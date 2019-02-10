package main

import (
	"fmt"
)

type MyNumberCard struct {
	id   int64
	name string
}

func main() {
	// sample := new(MyNumberCard)
	var sample *MyNumberCard = new(MyNumberCard)

	// All fields are zero value.
	// Zero value means
	//   numeric type: 0
	//   string type: empty string
	//   reference type: nil
	fmt.Println(sample)
	fmt.Println(sample.id)   // 0
	fmt.Println(sample.name) // empty sting

	me := new(MyNumberCard)
	me.id = 1
	me.name = "INSEO KIM"
	fmt.Println(me)
}
