/* Logical OR (||): Returns true if one of the statements is true. 
it means Only one statement should be true. if the both statements are false.
it will produce false.
*/

package main

import "fmt"

func main(){

	var a, b = 14, 25

	var e = a > b || a < b    
	fmt.Println(e)

	var f = a == 12 || a > b  // Output: False because both statements are false.
	fmt.Println(f)

}