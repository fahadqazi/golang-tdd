package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height, width float64
}

func (r Rectangle) Perimeter() float64 {
	return (2 * r.height) + (2 * r.width)
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type Triangle struct {
	height, base float64
}

func (t Triangle) Area() float64 {
	return t.height * t.base * 0.5
}
