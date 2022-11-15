package main

import (
	"fmt"
)

func main() {
	a := 35
	fmt.Printf("%d, %b , %#x\n", a, a, a)
	b := a << 1
	fmt.Printf("Left shift of a %v", b)

}
