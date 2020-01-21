package structs

import "math"

// Shape a
type Shape interface {
	Area() float64
}

// Rectangle a
type Rectangle struct {
	Width  float64
	Height float64
}

// Area for rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle a
type Circle struct {
	Radius float64
}

// Area for circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle a
type Triangle struct {
	Base   float64
	Height float64
}

// Area for rectangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter calculates
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}
