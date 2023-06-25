// Type Inference:
				/* if variable has a initial value. Go will automatically be able to infer 
				the type of that variable using that initial value. */

// Syntax:  var variable_name   = initialValue

package main

import "fmt"

func main(){

	var a = 0.5  

	fmt.Printf("type: %T\n", a)
}