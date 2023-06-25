/* How to Declare Multiple Variable:
	syntax: var name1, name2 type = initialvalue1, initialvalue2 
*/

package main

import "fmt"

func main() {

var name, b string = "Mr.", "Hello_World"
var x, y int = 25, 36

fmt.Println(name, b)
fmt.Println(x,y)

// type can be removed if the variable have initial Value
var n, c = "Mr." , "Hello_World"
fmt.Println("N: ",n, "\nC: ",c)

// if the variable has no initialValue. it will have assign 0 as a initialValue automatically

var p, q int
fmt.Println("P value is ", p , "\nQ value is", q)

}