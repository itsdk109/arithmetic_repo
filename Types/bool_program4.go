/* Logical NOT (!): Reverse the result, returns false if the result is true 

*/

package main

import "fmt"

func main() {
	a := 10

	c := !(a > 12 && a < 15)
	fmt.Println(c)  // output: True

	d := !(a == 10 && a < 12)
	fmt.Println(d)  // output: False

}