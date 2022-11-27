package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var waitgrp_obj sync.WaitGroup

	var x int64
	gs := 500
	waitgrp_obj.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&x, 1)
			y := atomic.LoadInt64(&x)

			fmt.Println(y)

			waitgrp_obj.Done()
		}()
	}
	waitgrp_obj.Wait()
	fmt.Println("final x:", x)
}
