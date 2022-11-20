package main

import (
	"fmt"
)

func main() {
	x := []int{23, 23, 432, 43, 55, 33, 33, 123}
	for index, value := range x {
		fmt.Println(index, value)
	}
	fmt.Printf("The type of x is %T", x)
}
