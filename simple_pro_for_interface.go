package main

import (
	"fmt"
	"math"
)

// create an interface
type MyInterface interface {
	area() float64 // area() is a method name
}

// create sturct for methods
type FincCicleArea struct {
	radius float64
}

type FindRectangleArea struct {
	width, height float64
}

// create a methods with func ()
func (rd FincCicleArea) area() float64 {

	return math.Pi * rd.radius * rd.radius
}

func (wh FindRectangleArea) area() float64 {

	return wh.width * wh.height
}

// accept interface

func ccir(cal MyInterface) {
	fmt.Printf("Area of Circle is %v\n", cal.area())
}

func cres(cal1 MyInterface) {
	fmt.Printf("Area of Rectangle is %v\n", cal1.area())
}

// main function

func main() {
	cr := FincCicleArea{5}
	cr1 := FindRectangleArea{10, 5}

	ccir(cr)
	cres(cr1)
}
