package main

import (
	"fmt"
)

func main() {
	sum_val := sum(33, 43)
	x, s := int_and_string()

	fmt.Println(sum_val, x, s)
}

func sum(a int, b int) int {
	return a + b
}

func int_and_string() (int, string) {
	return 1984, "Amey"
}
