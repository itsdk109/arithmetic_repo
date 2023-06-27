// Pointers

package main

import "fmt"

func main(){

	var a int = 10
	var p *int = &a
	fmt.Println(p)
	fmt.Println("Address of a is:", &a) // Output: print address of A 

	fmt.Println("Value of p is: ",p) // Output: it print Address of A variable.


	fmt.Println("P point value is",*p) // Output: it print Value of A variable.

}