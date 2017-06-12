package main

import (
	"fmt"
	"runtime"
	//"strconv"
	"sync"

	"golang.org/x/sync/syncmap"
)

func main() {
	maxproc := runtime.GOMAXPROCS(0)
	fmt.Printf("GOMAXPROCS=%d\n", maxproc)

	a := new(syncmap.Map)

	var wait sync.WaitGroup
	wait.Add(maxproc)
	for i := 0; i < maxproc; i++ {
		go func() {
			defer wait.Done()
			for i := 0; i < 1000000; i++ {
				//stri := strconv.Itoa(i)
				//a.Store(stri, stri)
				a.Store(i, i)
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
				//val, _ := a.Load(stri)
				val, _ := a.Load(i)
				_ = val
			}
		}()
	}
	wait.Wait()
}
