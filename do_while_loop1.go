// Do While Loop

package main

import "fmt"

func main() {

	var n int
	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &n)
	if n%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd number")
	}
}
