package main

import "fmt"

type person struct {
	first_name string
}

func changeMe(person_obj *person) {
	(person_obj).first_name = "Amit Joshi"
}
func main() {

	x := person{
		first_name: "Amey Dhongade",
	}
	fmt.Println(&x)
	changeMe(&x)
	fmt.Printf("New value : %v", x)
}
