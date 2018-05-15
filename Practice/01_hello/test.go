package main

import (
	"fmt"
)

type test struct {
	num1 int
	num2 int
}

func (t test) add() {
	fmt.Println(t.num1 + t.num2)
}
