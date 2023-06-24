// Arrays are used to store multiple values of the same type in a single variable.

package main 


import "fmt"

func main(){

	/* array declaration syntax : 
            var array_name = [length]datatype{values} 
	*/

	var a[5]int
	fmt.Println(a)   // output: [0,0,0,0,0]

	// to set array value on specific position
	a[2]=25
	fmt.Println("set: ",a)  // Output: set [0,0,25,0,0]

}