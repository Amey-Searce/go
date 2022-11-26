package main

import "fmt"

type person struct {
	first string
	last  string
	age   int16
}

func (perosn_obj person) speak() {
	fmt.Printf("My name is %v %v and my age is %v", perosn_obj.first, perosn_obj.last, perosn_obj.age)
}
func main() {
	person_obj := person{
		first: "Amey",
		last:  "Dhongade",
		age:   22,
	}
	person_obj.speak()
}
