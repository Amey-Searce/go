package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First_Name string `json:"First_name"`
	Last_Name  string `json:"Last_name"`
}

func main() {
	data_bytes := `[{"First_name":"Amey","Last_name":"Dhongade"},{"First_name":"Amey","Last_name":"Dhongade"}]`
	var person_obj []person
	err := json.Unmarshal([]byte(data_bytes), &person_obj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person_obj)

}
