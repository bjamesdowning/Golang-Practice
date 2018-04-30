package main

import (
	"bytes"
	"fmt"
)

/*
* Find the largest palindrome made from the product of two 3-digit numbers.
 */

//need to separate different actions.
func main() {
	var sum int
	var lrg int
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			sum = i * j
			str := fmt.Sprintf("%v", sum)
			rev := []byte{}
			forw := []byte{}
			//creates a reverse of the string from the origin sum
			for k := len(str) - 1; k > 0; k-- {
				rev = append(rev, str[k])
			}
			//turns str into a slive to compare to rev
			for l := 0; l < len(str)-1; l++ {
				forw = append(forw, str[l])
			}
			//use bytes package to compare slices
			if bytes.Compare(rev, forw) == 0 {
				if sum > lrg {
					lrg = sum
				}
			}
		}
	}
	fmt.Print("Largest palendrom is ", lrg, "\n")
}
