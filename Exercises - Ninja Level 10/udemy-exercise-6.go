package main

import "fmt"

func main() {
	channel_obj := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			channel_obj <- i
		}
		close(channel_obj)
	}()

	for values := range channel_obj {
		fmt.Println(values)
	}

}
