// How to create a struct

package main

import "fmt"

type Emp struct {
	Name   string
	Salary int
	Age    int
}

func main() {
	// access struct
	access_struct := Emp{Name: "Deepak",
		Salary: 158453,
		Age:    25,
	}
	fmt.Printf("%+v\n", access_struct)
}
