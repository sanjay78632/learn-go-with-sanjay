package shapes

import "math"

// ✅ Interface
type Shape interface {
	Area() float64
}

// ✅ Structs implementing the interface
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}