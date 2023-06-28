// Switch Statement 

package main

import "fmt"

func main(){
	var n int
	var num1, num2 int
	// code start here

	fmt.Print("Press 1:  Add |  \nPress 2:  Subtraction | \nPress 3:  Multiplication | \nPress 4:  Division | \nInput: ")
	fmt.Scanf("%d",&n)

	fmt.Println(" ")

	fmt.Print("Enter first Value: ")
	fmt.Scanf("%d",&num1)
	fmt.Print("Enter Second Value: ")
	fmt.Scanf("%d",&num2)

	switch (n) {
	case 1: add := num1 + num2 
	fmt.Println("Addition: ",add)
	break

	case 2: sub := num1 - num2
	fmt.Println("Sub: ",sub)
	break

	case 3: mul := num1 * num2
	fmt.Println("Multiplication: ",mul)
	break

	case 4: div := num1 / num2
	fmt.Println("Division: ",div)
	break

	default:
		fmt.Println("Enter Valid number")	
	}
}
