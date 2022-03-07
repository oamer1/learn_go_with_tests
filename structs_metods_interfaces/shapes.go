package main

import "math"

// Unlike Java and C++, Go does not allow function overloading
// func Area(circle Circle) float64       {}
// func Area(rectangle Rectangle) float64 {}

// so
// You can have functions with the same name declared in different packages. So we could create our Area(Circle) in a new package, but that feels overkill here.
// We can define methods on our newly defined types instead.

// A method is a function with a receiver

// Interfaces  are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different
// types and create highly-decoupled code whilst still maintaining type-safety
// In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}
type Shape interface {
	Area() float64
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

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
