// how to Anonymous Function within or outside the 'main' function

package main

import "fmt"

// create anonymous func outside the 'main' func

var outerFunction = func(x, y int) int {
	value1 := x + y
	return value1
}

func main() {

	// create anonymous func within the 'main' func
	func() {
		fmt.Println("Hello world!")
	}()
	// add program
	sum := func(x, y int) int {
		arg1 := x + y
		return arg1
	}(14, 7)
	fmt.Println(sum)
	fmt.Println(" ")
	/* ************************************** */
	// Find Area of Triangle
	var h, b int
	fmt.Print("Enter Height Value: ")
	fmt.Scanf("%d", &h)
	fmt.Print("Enter Based Value: ")
	fmt.Scanf("%d", &b)

	a := func(h, b int) int {
		area := (h * b) / 2
		return area
	}(h, b)

	fmt.Println("Area of Triangle: ", a)
	fmt.Println(" ")
	/* -------------------------------------------- */
	// create anonymous func outside the 'main' func
	/* calling function */
	add := outerFunction(14, 8)
	fmt.Println("Sum of two number is :", add)
}
