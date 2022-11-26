package main

import "fmt"

func sum_total(array_sum ...int) {
	total_sum := 0
	for _, value := range array_sum {
		total_sum += value
	}
	fmt.Println(total_sum)
}

func main() {

	array_numbers := []int{34, 35, 36}
	defer sum_total(array_numbers...)
	fmt.Println("Hello world is prinited before sum_total due to defer")
}
