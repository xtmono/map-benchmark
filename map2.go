package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	maxproc := runtime.GOMAXPROCS(0)
	fmt.Printf("default_max_procs=%d\n", maxproc)

	a := make(map[int]int)

	var wait sync.WaitGroup
	wait.Add(maxproc)
	for i := 0; i < maxproc; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < 1000000; i++ {
				a[i] = i
			}
		}()
	}
	wait.Wait()

	wait.Add(maxproc)
	for i := 0; i < maxproc; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < 1000000; i++ {
				val := a[i]
				_ = val
			}
		}()
	}
	wait.Wait()
}
