package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	sideLength float64
}

type square struct {
	base float64
	height float64
}


func main() {
	t := triangle{sideLength: 10}
	sq := square{base: 10, height: 10}

	printArea(t)
	printArea(sq)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.sideLength * t.sideLength
}

func (sq square) getArea() float64 {
	return sq.base * sq.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}