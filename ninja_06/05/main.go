package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

func (a square) area() float64 {
	area := a.length * a.length
	return area
}

func (a circle) area() float64 {
	area := a.radius * a.radius * math.Pi
	return area
}

func info(a shape) {
	fmt.Println(a.area())
}

func main() {
	a := square{
		length: 15,
	}
	b := circle{
		radius: 12.345,
	}
	info(a)
	info(b)
}
