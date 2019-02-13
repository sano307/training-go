package main

import (
	"fmt"
	"image"
)

// The color.Color and color.Model types are also interfaces,
// but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel.
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
