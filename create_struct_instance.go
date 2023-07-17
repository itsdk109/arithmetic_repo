package main

import "fmt"

type person struct {
	name    string
	address string
	phone   int
}

func main() {
	person_detail := person{"Deepak", "Bihar", 7854589535}
	fmt.Printf("%+v\n", person_detail)
	fmt.Println(" ")

	// create instance
	p := new(person)

	p.name = "Aryan"
	p.address = "Delhi"
	p.phone = 7854985635

	fmt.Printf("%+v\n", p)
}
