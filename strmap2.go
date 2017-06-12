package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

func main() {
	maxproc := runtime.GOMAXPROCS(0)
	fmt.Printf("default_max_procs=%d\n", maxproc)

	a := make(map[string]string)

	var wait sync.WaitGroup
	wait.Add(maxproc)
	for i := 0; i < maxproc; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < 1000000; i++ {
				stri := strconv.Itoa(i)
				a[stri] = stri
			}
		}()
	}
	wait.Wait()

	wait.Add(maxproc)
	for i := 0; i < maxproc; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < 1000000; i++ {
				stri := strconv.Itoa(i)
				val := a[stri]
				_ = val
			}
		}()
	}
	wait.Wait()
}
