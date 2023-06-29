// Slice

package main

import "fmt"

func main() {

	// Slice
	slice1 := []int{45, 21, 32}

	// get capacity of the slice
	fmt.Printf("Capacity of slice is %v\n", cap(slice1))

	// get length of the slice
	fmt.Printf("Length of slice is %v\n", len(slice1))

	// get value of the slice
	fmt.Printf("Slice value is %v\n", slice1)
}
