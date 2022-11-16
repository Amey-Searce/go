package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		val := i % 4
		fmt.Printf("Number : %d, Modulo 4 value:%d\n", i, val)
	}
}
