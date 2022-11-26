package main

import "fmt"

func greeting() func() string {
	greet_obj := func() string {
		return "Hello"
	}
	return greet_obj
}
func main() {
	f := greeting
	fmt.Printf("I want to say %v", f()())
}
