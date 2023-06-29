/* Range : The range keyword is used in for loop to iterate
over items of an array, slice, channel or map.
*/

package main

import "fmt"

func main() {

	// create array
	var arr = [5]int{1, 2, 3, 4, 5}

	// use range function in for loop
	for _, arrdata := range arr {
		fmt.Printf("Array Values => %d\n", arrdata)
	}

}
