package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Diameter float64
}

func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (d Circle) Area() float64 {
	return math.Pi * math.Pow((d.Diameter/2), 2)
}
