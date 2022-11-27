package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First_name string
	Last_name  string
}

func main() {

	person_obj := person{
		First_name: "Amey",
		Last_name:  "Dhongade",
	}
	person_obj_1 := person{
		First_name: "Suresh",
		Last_name:  "Dhongade",
	}
	marshal_objects := []person{person_obj, person_obj_1}

	marshal_obj, err := json.Marshal(marshal_objects)
	if err != nil {
		fmt.Printf("Error is %v", err)
	}
	fmt.Println(string(marshal_obj))
}
