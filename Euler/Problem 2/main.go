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
	xs := []int{1, 2}
	ans := 2
	for {
		fib := xs[1] + xs[0]
		xs = append(xs, fib)
		xs = append(xs[:0], xs[1:]...) //Avoid expanding slice beyond underlying arr
		if fib > 4000000 {
			break
		} else if fib%2 == 0 {
			ans += fib
		}
	}
	fmt.Println(ans)

}
