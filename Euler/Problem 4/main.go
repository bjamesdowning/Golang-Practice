package main

import (
	"fmt"
)

/*
* Find the largest palindrome made from the product of two 3-digit numbers.
 */

//Not a very good solution. Good enough for now.
func main() {
	var sum int
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
			if rev[0] == forw[0] && rev[1] == forw[1] && rev[2] == forw[2] {
				fmt.Print("Largest palendrom is ", sum)
				fmt.Println("----", i, j)
				break
			}
		}
	}
}
