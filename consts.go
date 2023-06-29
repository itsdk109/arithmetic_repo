package main

import "fmt"

func main() {
	const a = 10
	fmt.Println(a)

	// trying to change const value
	a = 45
	fmt.Println(a) // error: ./consts.go:9:3: cannot assign to a (untyped int constant 10)
}
