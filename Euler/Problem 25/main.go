package main

import "fmt"

func fib() int {
	xs := []int{0, 1}
	for i := 1; i > 0; i++ {
		fib := xs[i] + xs[i-1]
		xs = append(xs, fib)
		fmt.Println(rune(fib))
		if fib > 100 {
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(fib())
}
