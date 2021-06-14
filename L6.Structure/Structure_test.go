package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{20.0, 10.0}
	got := Perimeter(rectangle.Width, rectangle.Height)
	want := 60.0

	if got != want {
		t.Errorf("want: '%v', got '%v'", want, got)
	}
}

func TestArea(t *testing.T) {

	t.Run("Area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{20.0, 10.0}
		got := rectangle.Area()
		want := 200.0
		if got != want {
			t.Errorf("want: '%.2f', got '%.2f'", want, got) //float64 -> %.2f
			// https://golang.org/pkg/fmt/
			// %.2f   default width, precision 2
		}
	})

	t.Run("Area of a circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 25 * math.Pi
		if got != want {
			t.Errorf("want: '%.2f', got '%.2f'", want, got)
		}
	})

}
