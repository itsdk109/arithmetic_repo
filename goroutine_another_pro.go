package main

import (
	"fmt"
	"math"
	"time"
)

func Calculator(num1, num2 int) {
	var slice = []int{1, 3, 5, 8, 12, 15, 22, 42}

	for _, slice_value := range slice {
		if slice_value == 3 {
			sum := num1 + num2
			fmt.Printf("Sum of Two Number %v + %v = %v\n", num1, num2, sum)
		} else if slice_value == 8 {
			sub := num1 - num2
			fmt.Printf("Subtraction of two number %v - %v = %v\n", num1, num2, sub)
		} else if slice_value == 15 {
			mul := num1 * num2
			fmt.Printf("Multiply of two number %v * %v = %v\n", num1, num2, mul)
		} else if slice_value == 42 {
			div := num1 / num2
			fmt.Printf("Divide of two number %v/%v = %v\n", num1, num2, div)
		}
	}
}

func Mensuration() {
	areaofRectangle := func(width, height float64) float64 {

		return width * height
	}

	areaofcircle := func(r float64) float64 {

		return math.Pi * r * r
	}

	width := 10.12
	height := 12.14
	fmt.Printf("Area of rectangle %v * %v = %v\n", width, height, areaofRectangle(width, height))

	r := 2.8
	fmt.Printf("Area of Circle %v * %v* %v = %v\n", math.Pi, r, r, areaofcircle(r))
}

func main() {

	go Calculator(14, 10)
	time.Sleep(1 * time.Second)
	fmt.Println(" ")
	fmt.Println("<<This is mensuration program>>\n")
	go Mensuration()
	time.Sleep(1 * time.Second)
}
