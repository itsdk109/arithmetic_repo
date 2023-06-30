/* write a program to print sum of
user entered number usign  for loop */

package main

import "fmt"

func main() {

	var n int
	var sum = 0

	// take input from user
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	fmt.Println("Sum of number is : ", sum)
}
