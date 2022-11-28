package main

import (
	"exercises/Exercises-NinjaLevel12/dog"
	"fmt"
)

type dog1 struct {
	name string
	age  int
}

func main() {

	dog_obj := dog1{
		name: "Bolt",
		age:  dog.Years(23),
	}
	fmt.Println(dog_obj.age)
}
