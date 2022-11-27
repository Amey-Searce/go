package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("No of cpus", runtime.NumCPU())
	fmt.Println("No of routines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("First routine starts")
		wg.Done()
	}()

	go func() {
		fmt.Println("Second routine starts")
		wg.Done()
	}()

	fmt.Println("No of cpus", runtime.NumCPU())
	fmt.Println("No of routines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("After the routines are completed")
	fmt.Println("No of cpus", runtime.NumCPU())
	fmt.Println("No of routines", runtime.NumGoroutine())
}
