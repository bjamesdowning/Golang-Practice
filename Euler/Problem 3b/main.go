package main

import "fmt"

//Faster and simpler solution to Euler problem 3.
func main() {
	fmt.Println(findPrime(600851475143, 2))
}

//Use of recursive function
func findPrime(a int64, b int64) int64 {
	if a%b == 0 {
		a /= b
		return findPrime(a, b)
	} else if b >= a {
		return b
	} else {
		b++
		return findPrime(a, b)
	}
}
