// Write a program to find the table from given array using range function.

package main

import "fmt"

func main() {

	// create an Array

	var arr_data = []int{2, 5, 6, 7, 8}

	for _, table := range arr_data { // [14,52,66,70,68]
		if table == 6 {
			for i := 1; i <= 10; i++ {
				table := table * i
				fmt.Println(table)

			}
		} else {
			fmt.Println("")
		}
	}
}
