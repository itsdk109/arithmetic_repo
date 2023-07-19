package main

import "fmt"

type Employee struct {
	Name    string
	ID      int
	Salary  int
	Mobile  int
	Address string
}

/*
emp Employee --> receivertype
emp_detail() --> method name
Employee (Struct Name) --> returntype
*/

// create a method
func (emp Employee) emp_detail() Employee {

	return emp

}

func main() {
	show_detail := Employee{"Deepak", 1, 78542, 785425895, "Buxar"}

	fmt.Println("EmPloyee Details: ", show_detail.emp_detail())
}
