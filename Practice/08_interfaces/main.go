package main

import (
	"fmt"
	"math"
)

//Square struct
type Square struct {
	side float64
}

//Square method on type Square
func (s Square) area() float64 {
	return s.side * s.side
}

//Circle struct
type Circle struct {
	radius float64
}

//Circle method on type Circle
func (r Circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

//Shape interface
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	sq := Square{10}
	crl := Circle{5}
	info(sq)
	info(crl)
}
