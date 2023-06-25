/*
	How to create different type variable in single statement.

	syntax: var (
					name = "Mr. Go Language"
					age = 16
				)

*/

package main 

import "fmt"

func main() {
	var (
		name = "Hello World"
		age = 25
		boolean = true
	)

	fmt.Println("Name: ", name, "\nAge: ", age, "\nBoolean Value is: ",boolean)

}