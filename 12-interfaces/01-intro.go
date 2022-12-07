package main

import (
	"fmt"
	"math"
)

//3rd party
//********************
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

//********************

type AreaCalculator interface {
	Area() float32
}

//utility functions
func PrintArea(x AreaCalculator) {
	fmt.Println("Area = ", x.Area())
}

type PerimeterCalculator interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterCalculator) {
	fmt.Println("Perimeter = ", x.Perimeter())
}

type ShapeCalculator interface {
	AreaCalculator
	PerimeterCalculator
}

func PrintShape(x ShapeCalculator) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{Radius: 12}
	// fmt.Println("Area = ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 10, Width: 12}
	// fmt.Println("Area = ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
