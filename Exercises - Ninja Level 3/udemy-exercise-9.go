package main

import (
	"fmt"
)

func main() {

	var favSport string = "football"
	switch favSport {
	case "cricket":
		{
			fmt.Printf("Cricket")
		}
	case "football":
		{
			fmt.Printf("Football")
		}
	}
}
