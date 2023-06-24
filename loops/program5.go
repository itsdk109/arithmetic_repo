/* To Print this pattern -
		*******
		******
		*****
		****
		***
		**
		*
*/

// Main Code Start

package main

import "fmt"

func main() {

	n := 7
	
	for i := 6 ; i < n ; i-- {

		for j := 0 ; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}