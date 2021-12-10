package main

import (
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

type Triangle struct {
	side float32
	hight float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (tr *Triangle) Area() float32 {
	return tr.side * tr.hight / 2
}

func main() {
	sq1 := new(Square)
	sq1.side = 5
	fmt.Printf("The square has area: %f\n", sq1.Area())

	tr1 := new(Triangle)
	tr1.side = 5
	tr1.hight = 8
	fmt.Printf("The triangle has area: %f\n", tr1.Area())
}