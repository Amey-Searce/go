package main

import (
	"fmt"
)

func main() {
	x := [4]int{23, 23, 432, 43}
	for index, value := range x {
		fmt.Println(index, value)
	}
	fmt.Printf("The type of x is %T", x)
}
