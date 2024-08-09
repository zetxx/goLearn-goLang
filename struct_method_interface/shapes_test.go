package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.00, 10.00}
	got := rectangle.Perimeter()
	want := 40.00

	if got != want {
		t.Errorf("got %.2f != want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	shapes := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2.00, 2.00}, 4.00},
		{Circle{10.00}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, d := range shapes {
		got := d.shape.Area()
		if got != d.want {
			t.Errorf("got %.2f != want %.2f", got, d.want)
		}
	}
}
