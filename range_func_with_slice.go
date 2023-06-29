package main

import "fmt"

func main() {

	// create slice
	var slice_data = []int{14, 15, 16, 17, 18}

	for _, sdata := range slice_data {
		fmt.Println(sdata)
	}

	fmt.Println(" ")

	// without index parameter
	for sdata := range slice_data {
		fmt.Println(slice_data[sdata])
	}
}
