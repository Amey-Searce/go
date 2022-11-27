package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{123, 35, 67}
	string_arr := []string{"Amey", "Dhongade"}
	sort.Ints(arr)
	sort.Strings(string_arr)
	fmt.Println(arr)
	fmt.Println(string_arr)
}
