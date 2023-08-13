package main

import (
	"fmt"
)

func printTable(number int, table chan int) {
	for i := 1; i <= 10; i++ {
		table <- number * i
	}
	close(table)
}

func main() {
	var number int
	// take input from user
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &number)

	// create a channel
	cha := make(chan int)

	// call the function with goroutine
	go printTable(number, cha)

	fmt.Printf("table of %d\n", number)
	for result := range cha {
		fmt.Println(result)
	}
}
