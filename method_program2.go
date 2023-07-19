// How to create Multiple Method using a struct

package main

import "fmt"

// create a struct
type calculator struct {
	num1, num2 int
}

func (c calculator) sum() int {

	return c.num1 + c.num2
}

func (c calculator) sub() int {

	return c.num1 - c.num2
}

func (c calculator) mul() int {

	return c.num1 * c.num2
}

func (c calculator) div() int {

	return c.num1 / c.num2
}

func main() {
	call_calculator := calculator{44, 5}
	fmt.Println("Sum of two number is: ", call_calculator.sum())
	fmt.Println("Subtract of two number is: ", call_calculator.sub())
	fmt.Println("Multiply of two number is: ", call_calculator.mul())
	fmt.Println("Division of two number is: ", call_calculator.div())
}
