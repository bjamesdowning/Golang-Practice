package main

import (
	"bytes"
	"fmt"
)

/*
* Find the largest palindrome made from the product of two 3-digit numbers.
 */

//Not a very good solution. Good enough for now.
func main() {
	var sum int
	var lrg int
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			sum = i * j
			str := fmt.Sprintf("%v", sum)
			rev := []byte{}
			forw := []byte{}
			for k := len(str) - 1; k > 0; k-- {
				rev = append(rev, str[k])
			}
			for l := 0; l < len(str)-1; l++ {
				forw = append(forw, str[l])
			}
			if bytes.Compare(rev, forw) == 0 {
				if sum > lrg {
					lrg = sum
				}
			}
		}
	}
	fmt.Print("Largest palendrom is ", lrg, "\n")
}
