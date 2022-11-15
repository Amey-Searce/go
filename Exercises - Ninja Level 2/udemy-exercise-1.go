package main

import (
	"fmt"
)

func main() {
	var number int = 43
	fmt.Printf("Number in integer format %d\n", number)
	fmt.Printf("Number in binary format %b\n", number)
	fmt.Printf("Number in hexadecimal format %#x\n", number)
}
