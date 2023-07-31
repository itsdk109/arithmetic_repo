// else if program

package main

import "fmt"

func add(x, y int) int {
	// create a variable
	var num1, num2 int

	if x >= 1 && x <= 10 {
		num1 = x

	} else {
		fmt.Println("Error: Please Enter Number of X within 1 and 10")
		return 0
	}

	if y >= 10 && y <= 20 {
		num2 = y

	} else {
		fmt.Println("Error: Please Enter Number of Y within 10 and 20")
		return 0
	}
	var sum = num1 + num2
	return sum
}

func main() {
	var x, y int
	fmt.Print("#Kindly Enter number of X Within 1-10 & Y within 10-20")
	fmt.Println(" ")
	fmt.Print("Enter X value: ")
	fmt.Scanf("%d", &x)

	fmt.Print("Enter Y value: ")
	fmt.Scanf("%d", &y)

	var result int = add(x, y)
	fmt.Println(result)
}
