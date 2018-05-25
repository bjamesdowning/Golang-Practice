package main

import (
	"flag"
	"fmt"
)

//Find the sum of the digits in the number 100!
func main() {
	f := flag.Int("n", 0, "Integer")
	flag.Parse()
	fmt.Println(getFactorial(*f))
}

func getFactorial(n int) int {
	for i := n - 1; i > 0; i-- {
		n *= i
	}
	return n
}
