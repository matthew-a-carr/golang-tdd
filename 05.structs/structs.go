package structs

import "math"

type Rectange struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectange) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r Rectange) Area() float64 {
	return (r.width * r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
