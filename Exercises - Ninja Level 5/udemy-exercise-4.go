package main

import "fmt"

func main() {
	s := struct {
		first_name string
		last_name  string
	}{
		first_name: "Rahul",
		last_name:  "Dravid",
	}
	fmt.Println(s.first_name)
}
