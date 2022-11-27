package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var waitgrp_obj sync.WaitGroup

	x := 0
	gs := 500
	waitgrp_obj.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := x
			runtime.Gosched()
			v++
			x = v
			fmt.Println(x)
			waitgrp_obj.Done()
		}()
	}
	waitgrp_obj.Wait()
	fmt.Println("final x:", x)
}
