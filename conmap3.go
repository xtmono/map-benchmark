package main

import (
	"fmt"
	"runtime"
	//"strconv"
	"sync"

	"github.com/fanliao/go-concurrentMap"
)

func main() {
	maxproc := runtime.GOMAXPROCS(0)
	fmt.Printf("GOMAXPROCS=%d\n", maxproc)

	a := concurrent.NewConcurrentMap()

	var wait sync.WaitGroup
	wait.Add(maxproc)
	for i := 0; i < maxproc; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < 1000000; i++ {
				//stri := strconv.Itoa(i)
				//a.Put(stri, stri)
				a.Put(i, i)
			}
		}()
	}
	wait.Wait()

	wait.Add(maxproc)
	for i := 0; i < maxproc; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < 1000000; i++ {
				//stri := strconv.Itoa(i)
				//val, _ := a.Get(stri)
				val, _ := a.Get(i)
				_ = val
			}
		}()
	}
	wait.Wait()
}
