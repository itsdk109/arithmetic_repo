/* How to invoke closure function withing main function */

package main

import "fmt"

// create closure function
func calculator() func(b int) (int, int, int, int) {
	var a int
	// get input from the user
	fmt.Println("Enter First Number: ")
	fmt.Scanf("%d", &a)

	return func(b int) (int, int, int, int) {
		sum := a + b
		sub := a - b
		mul := a * b
		div := a / b // ('/')--> it returns quotient(Bhagfal) value

		return sum, sub, mul, div
	}
}

func areaOfSquare() func(a int) int {

	return func(a int) int {
		area := a * a
		return area
	}
}

func main() {

	// invoke calculator function
	calling_func := calculator()

	var b int
	fmt.Print("Enter Second Number: ")
	fmt.Scanf("%d", &b)
	fmt.Println(" ")
	add, sub, mul, div := calling_func(b)
	fmt.Println("Sum of Two Number is: ", add, " ")
	fmt.Println("Subtraction of Two Number is: ", sub, " ")
	fmt.Println("Product of Two Number is: ", mul, " ")
	fmt.Println("Division of Two Number is: ", div, " ")

	fmt.Println(" ")

	// area of Square
	call_area := areaOfSquare()

	var a int
	fmt.Print("Enter Side value of square: ")
	fmt.Scanf("%d", &a)
	result := call_area(a)
	fmt.Printf("Area of Square: %d * %d = %d\n ", a, a, result)

}
