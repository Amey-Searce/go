package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (square_obj square) area() float64 {
	return square_obj.length * square_obj.length
}
func (circle_obj circle) area() float64 {
	return math.Pi * circle_obj.radius * circle_obj.radius
}

type shape interface {
	area() float64
}

func info(shape_obj shape) {
	fmt.Println(shape_obj.area())
}
func main() {

	circle_obj := circle{
		radius: 34.5,
	}
	square_obj := square{
		length: 13.45,
	}
	info(circle_obj)
	info(square_obj)

}
