package main

import (
	"fmt"
)

const (
	pi         = 3.14
	g  float32 = 9.82
)

func main() {
	fmt.Printf("Area of a circle radius 2 %v\n", 2*pi*2)
	fmt.Printf("Force with 10 kg mass %v", 10*9.82)

}
