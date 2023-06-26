/* Logical AND : Returns true if both statements are true. 
it means both statements should be True otherwise it will produce false value. */

package main

import "fmt"

func main() {
	a := 10
	b := a < 12 && a > 4

	fmt.Println(b)  // output: True becasue both statements are true.

	c := a < 5 && a == 10 // Output: False because one statement is true.

	fmt.Println(c)


}