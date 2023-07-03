// How to declare 2D Array In Go Language

/* Syntax:
	var variable_name = [rows_size][column_size] datatype {
					{0,0,0},
					{0,0,0},
					{0,0,0},
}
*/

package main

import "fmt"

func main() {
	// create 2D Array
	var arr_data = [4][3]int{{4, 2, 3}, {3, 4, 5}, {5, 6, 10}, {7, 8, 19}}

	for i := 0; i < 4; i++ {

		for j := 0; j < 3; j++ {

			fmt.Print(arr_data[i][j], "\t")

		}
		fmt.Println(" ")
	}

}
