package main

import "fmt"

type person struct {
	first_name       string
	last_name        string
	ice_cream_flavor []string
}

func main() {
	x := person{
		first_name:       "Sachin",
		last_name:        "Tednulkar",
		ice_cream_flavor: []string{"Vanilla", "Mix"},
	}
	y := person{
		first_name:       "Rahul",
		last_name:        "Dravid",
		ice_cream_flavor: []string{"Vanilla", "Mix"},
	}

	map_obj := map[string]person{
		x.last_name: x,
		y.last_name: y,
	}
	for index, value := range map_obj {
		fmt.Println(index, value)
		for index2, value2 := range value.ice_cream_flavor {
			fmt.Println(index2, value2)
		}
	}

}
