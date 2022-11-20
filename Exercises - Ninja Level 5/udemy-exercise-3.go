package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {

	x := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourwheel: true,
	}
	y := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}
	fmt.Println(y)
	fmt.Println(x.vehicle.color)
	fmt.Println(x.vehicle.doors)
}
