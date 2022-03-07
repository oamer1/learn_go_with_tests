package main

import "math"

// Unlike Java and C++, Go does not allow function overloading
// func Area(circle Circle) float64       {}
// func Area(rectangle Rectangle) float64 {}

// so
// You can have functions with the same name declared in different packages. So we could create our Area(Circle) in a new package, but that feels overkill here.
// We can define methods on our newly defined types instead.

// A method is a function with a receiver

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
