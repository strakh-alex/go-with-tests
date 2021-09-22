package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Area return area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area return area of Circle
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter returns a perimeter of Rectangle
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// Perimeter returns a perimeter of Circle
func (c Circle) Perimeter() float64 {
	return c.Radius * (math.Pi * 2)
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height + math.Sqrt(math.Pow(t.Base, 2) + math.Pow(t.Height, 2))
}