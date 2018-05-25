package main

import (
	"fmt"
)

//Use of labeling for loops then using labels to break
//Use of switch in place of multiple if's
func main() {
	var (
		evenIf  int
		oddIf   int
		totalIf int
		zerosIf int
		evenSw  int
		oddSw   int
		totalSw int
		zerosSw int
	)

	numbers := []int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10}
Abort:
	for i := 0; i < 10; i++ {
		for _, n := range numbers {
			totalIf++
			if n == 0 {
				zerosIf++
				break Abort
			}
			if n%2 == 0 {
				evenIf++
			} else {
				oddIf++
			}
		}
	}
	fmt.Printf("IF:  Odd numbers: %d, Even numbers: %d, Total numbers: %d\n", oddIf, evenIf, totalIf)

	//switch example of the same above code
	//
	for _, n := range numbers {
		totalSw++
		switch {
		case n == 0:
			zerosSw++
		case n%2 == 0:
			evenSw++
		default:
			oddSw++
		}
	}
	fmt.Printf("SWITCH:  Odd numbers: %d, Even numbers: %d, Zeroes %d, Total numbers: %d\n", oddSw, evenSw, zerosSw, totalSw)

}
