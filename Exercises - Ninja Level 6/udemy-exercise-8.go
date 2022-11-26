package main

import "fmt"

func sum() func() string {

	return func() string {
		return "Hello"
	}
}
func main() {

	obj := sum()
	fmt.Println(obj())

}
