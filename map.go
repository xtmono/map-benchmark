package main

func main() {
	a := make(map[int]int)

	for i := 0; i < 1000000; i++ {
		a[i] = i
	}

	for i := 0; i < 1000000; i++ {
		val := a[i]
		_ = val
	}
}
