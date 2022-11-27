package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitgrp_obj sync.WaitGroup

	x := 0
	gs := 500
	waitgrp_obj.Add(gs)
	var lock_obj sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			lock_obj.Lock()
			v := x
			// runtime.Gosched()
			v++
			x = v
			fmt.Println(x)
			lock_obj.Unlock()
			waitgrp_obj.Done()
		}()
	}
	waitgrp_obj.Wait()
	fmt.Println("final x:", x)
}
