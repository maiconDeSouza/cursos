package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

type cicle struct {
	radius float64
}

func (c cicle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type Forma interface {
	area() float64
}

func calculateArea(f Forma) {
	fmt.Println("A area é", f.area())
}

func main() {
	r := rectangle{length: 3, width: 4}
	c := cicle{radius: 12}

	calculateArea(r)
	calculateArea(c)
}
