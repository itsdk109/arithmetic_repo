//Display n terms of natural numbers and their sum
package main 

import "fmt"

func main() {
var sum int = 0
var numbers int 
fmt.Print("Enter the number: ")
fmt.Scan(&numbers)

for i := 1; i<=numbers; i++ {
	fmt.Println(i)
	sum = sum + i
}
	fmt.Print("Total Sum of N terms of natural numbers : ", sum)
}