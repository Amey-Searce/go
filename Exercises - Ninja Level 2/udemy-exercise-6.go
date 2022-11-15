package main

import (
	"fmt"
)

const (
	first_year  = 2017 + iota
	second_year = 2017 + iota
	third_year  = 2017 + iota
	fourth_year = 2017 + iota
)

func main() {
	fmt.Println(first_year, second_year, third_year, fourth_year)
}
