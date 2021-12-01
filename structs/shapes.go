package main

import "math"

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func Perimeter(rectangle Rectangle) float64 {
	sum := 2.0 * (rectangle.Length + rectangle.Width)
	return sum
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Length * rectangle.Width
// }

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

type Shape interface {
	Area() float64
}
