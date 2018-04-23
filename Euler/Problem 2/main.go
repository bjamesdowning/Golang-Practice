package main

import (
	"fmt"
)

/*
* Solve Euler Problem 2:By considering the terms in the
* Fibonacci sequence whose values do not exceed four million,
* find the sum of the even-valued terms.
 */
func main() {
	fmt.Println(evenFib(4000000))
}
func evenFib(limit int) int {
	xs := []int{1, 2}
	ans := 2
	for {
		fib := xs[1] + xs[0]
		xs = append(xs, fib)
		xs = append(xs[:0], xs[1:]...) //Avoid expanding slice beyond underlying arr
		if fib > limit {
			break
		} else if fib%2 == 0 {
			ans += fib
		}
	}
	return ans
}
