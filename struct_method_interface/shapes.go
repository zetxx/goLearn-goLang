package main

import "math"

type Rectangle struct {
	w float64
	h float64
}

type Triangle struct {
	base float64
	h    float64
}

type Circle struct {
	r float64
}

type Shape interface {
	Area() float64
}

func (dimensions Rectangle) Perimeter() float64 {
	return (dimensions.h + dimensions.w) * 2
}

func (dimensions Rectangle) Area() float64 {
	return dimensions.h * dimensions.w
}

func (dimensions Circle) Area() float64 {
	return dimensions.r * dimensions.r * math.Pi
}

func (dimensions Triangle) Area() float64 {
	return dimensions.base * dimensions.h * 0.5
}
