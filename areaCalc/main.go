package main

import (
	"fmt"
)

type square struct {
	side float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() (float64, string)
}

func main() {
	sq := square{side: 5}
	tr := triangle{height: 5, base: 3}
	printArea(sq)
	printArea(tr)
}

func (s square) getArea() (float64, string) {
	return s.side * s.side, "square"
}

func (t triangle) getArea() (float64, string) {
	return .5 * (t.height * t.base), "triangle"
}

func printArea(s shape) {
	a, sh := s.getArea()
	msg := "the area of this " + sh + " is %v units. \n"
	fmt.Printf(msg, a)
}
