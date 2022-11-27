package main

import "fmt"

type person struct {
	first_name string
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println("Amey Dhongade")
}

func saySomething(human_obj human) {
	human_obj.speak()
}

func main() {
	person_obj := person{
		first_name: "Amey Dhongade",
	}

	// if we have to use pointer change the function to *
	saySomething(person_obj)

	person_obj.speak()
}
