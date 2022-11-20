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

	fmt.Println(x.first_name, x.last_name, x.ice_cream_flavor)
	for index, value := range x.ice_cream_flavor {
		fmt.Println(index, value)
	}
	for index, value := range y.ice_cream_flavor {
		fmt.Println(index, value)
	}
	fmt.Println(y.first_name, y.last_name, y.ice_cream_flavor)

}
