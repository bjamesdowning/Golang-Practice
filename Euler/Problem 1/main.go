package main

import (
	"fmt"
)

//Solve Euler Problem 1. Find the sum of all multiples of 3 or 5 up to 1000
func main() {
	var sum int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
