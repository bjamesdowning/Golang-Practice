package main

import (
	"fmt"
)

//Solve Euler Problem 1. Find the sum of all multiples of 3 or 5 up to 1000
func main() {
	var arr []int
	var sum int
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			arr = append(arr, i)
		}
	}
	for _, v := range arr {
		sum += v
	}
	//fmt.Println(arr)
	fmt.Println(sum)
}
