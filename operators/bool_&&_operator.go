// Bool : it holds value with two states : True or False.
// && Operator: it returns only True when both a and b are true.


package main 

import "fmt"

func main() {
	a := true
	b := false
	c := a && b
	fmt.Println(c)  // it returns false 
}