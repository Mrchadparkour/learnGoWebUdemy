package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func main() {
	s1 := square{20}
	c1 := circle{4}

	getArea(s1)
	getArea(c1)
}

func (s square) area() {
	area := s.side * s.side
	fmt.Println("Area of square:", area)
}

type circle struct {
	radius float64
}

func (c circle) area() {
	area := math.Pi * c.radius * c.radius
	fmt.Println("Area of circle:", area)
}

type shape interface {
	area()
}

func getArea(s shape) {
	s.area()
}
