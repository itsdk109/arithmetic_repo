// How to find sum, sub, mul, div using interface

package main

import "fmt"

// create an interface

type CalInterface interface {
	// define methods
	sum() float64
	sub() float64
	mul() float64
	div() float64
}

// create a sturct for access methods

type Calculator struct {
	num1, num2 float64
}

// create a methods

func (c Calculator) sum() float64 {

	return c.num1 + c.num2
}

func (c Calculator) sub() float64 {
	return c.num1 - c.num2
}

func (c Calculator) mul() float64 {
	return c.num1 * c.num2
}

func (c Calculator) div() float64 {
	return c.num1 / c.num2
}

// access interface

func AcceptInterface(acci CalInterface) {
	fmt.Printf("Sum of two number  is %v\n", acci.sum())
	fmt.Printf("Subtract of two number is %v\n", acci.sub())
	fmt.Printf("Multiply of two number is %v\n", acci.mul())
	fmt.Printf("Divide of two number is %v\n", acci.div())
}

func main() {
	res := Calculator{12.5, 54.3}
	AcceptInterface(res)
}
