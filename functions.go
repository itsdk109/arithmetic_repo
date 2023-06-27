// How to declare function in go

/*
	Syntax: 
	func functionname(parametername type) returntype {  
 		//function body
	}
*/

package main
import "fmt"

func add(x int, y int) int {
	return x+y
}

func main(){

	x := 14
	y := 25

	sum := add(x,y)
	fmt.Println(sum)

}