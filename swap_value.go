// This is Swap Program where you can swap value.

// For Example: a = 10 , b= 20 ---->After Swaping----> a = 20 , b = 10

package main
import "fmt"

func main(){

	a := 10
	b := 20

	fmt.Println("A:",a," ","B:",b) 

	// Swap code
	a,b = b,a
	fmt.Println(" ")
	fmt.Println("After Swaping: \nA value is:",a, " ","B value is:",b )
}