package main

import (
	"fmt"
)

func table(n int) {
	for i := 1; i <= 10; i++ {
		table := i * n
		fmt.Printf("%d * %d = %d\n", n, i, table)
	}
}

func main() {

	// take input from the user
	var n int
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &n)

	table(n)

}
