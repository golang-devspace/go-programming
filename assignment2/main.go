package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	fmt.Println("Assignment 2 works...")
	ts := triangle{
		base:   10,
		height: 10,
	}

	ss := square{
		sideLength: 10,
	}

	printArea(ts)
	printArea(ss)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
