package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("For ascii value %d, Rune character:\n", i)
		for k := 1; k <= 3; k++ {
			fmt.Printf("%U\n", i)
		}
	}
}
