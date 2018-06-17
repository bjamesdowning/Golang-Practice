package main

import (
	"flag"
	"fmt"
	"strconv"
)

//Find the sum of the digits in the number 100!
func main() {
	f := flag.Int("n", 0, "Integer")
	flag.Parse()
	addInts(getFactorial(*f))
}

func getFactorial(n int) int {
	for i := n - 1; i > 0; i-- {
		n *= i
	}
	return n
}

func addInts(n int) {
	var res int
	itoS := strconv.Itoa(n)
	for _, i := range itoS {
		ni, err := strconv.Atoi(string(i))
		if err != nil {
			fmt.Println("ERROR ERROR:", err)
		}
		res += ni
	}
	fmt.Println(res)
}
