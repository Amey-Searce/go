package main

import "fmt"

func main() {
	channel_obj := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				channel_obj <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-channel_obj)
	}

}
