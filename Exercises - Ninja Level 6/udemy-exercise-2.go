package main

import "fmt"

func sum_total(sum ...int) int {
	total_sum := 0
	for _, value := range sum {
		total_sum += value
	}
	return total_sum
}
func total_sum(sum []int) int {
	total_sum := 0
	for _, value := range sum {
		total_sum += value
	}
	return total_sum
}

func main() {
	array_numbers := []int{34, 34, 34, 35}
	sum_returned := sum_total(array_numbers...)
	total_sum_returned := total_sum(array_numbers)
	fmt.Printf("Sum returned from both functions %v %v", sum_returned, total_sum_returned)
	
}
