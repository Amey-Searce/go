package main

import (
	"fmt"
)

func main() {
	x := 45
	if x == 45 {
		fmt.Printf("Numbers are equal")
	} else if x >= 45 {
		fmt.Printf("Number is greater than 45")
	} else {
		fmt.Printf("Number is smaller than 45")
	}
}
