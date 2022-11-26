package main

import "fmt"

func area(area func(area_obj float64) float64) {
	fmt.Println(area(3))
}
func main() {
	g := func(radius float64) float64 {
		return 3.14 * radius * radius
	}
	area(g)
}
