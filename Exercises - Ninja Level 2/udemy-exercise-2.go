package main

import (
	"fmt"
)

func main() {
	a := (136 == 136)
	b := (345 >= 123)
	c := (123 <= 345)
	d := (1 != 1)
	e := (3 > 4)
	f := (-1 < -100)
	fmt.Printf("Comaparison values:%v,%v,%v,%v,%v,%v", a, b, c, d, e, f)
}
