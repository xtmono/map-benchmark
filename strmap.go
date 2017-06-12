package main

import (
	"strconv"
)

func main() {
	a := make(map[string]string)

	for i := 0; i < 1000000; i++ {
		stri := strconv.Itoa(i)
		a[stri] = stri
	}

	for i := 0; i < 1000000; i++ {
		stri := strconv.Itoa(i)
		val := a[stri]
		_ = val
	}
}
