package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi this is anonymous")
	}()
}
