package main

import (
	"fmt"
)

func main() {
	c := gen()
	retrieve_values(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
func retrieve_values(channel_obj <-chan int) {
	for value := range channel_obj {
		fmt.Println(value)
	}
}
