// Switch Statement Program

package main

import "fmt"

func areaoftriangle(h, b int) int {
	var area = (h * b) / 2
	return area
}

func rectangle_perimeter(l, w int) int {
	var rec_perimeter = 2 * (l + w)
	return rec_perimeter
}

func print_table(p int) int {

	for i := 1; i <= 10; i++ {
		p := p * i
		fmt.Println(p)
	}
	return 0
}

// main function code

func main() {

	var n int
	fmt.Print("Press:\n1. Area of Triangle\n2. Rectangle Perimeter\n3.Print Table\n => ")
	fmt.Scanf("%d", &n)

	switch n {

	case 1:
		/* code */
		areoft := areaoftriangle(14, 7)
		fmt.Println("Area of Triangle: ", areoft)

	case 2:
		/* code */
		rect_perimeter := rectangle_perimeter(12, 3)
		fmt.Println("Rectangle Perimeter: ", rect_perimeter)

	case 3:
		/* code */
		pt := print_table(4)
		fmt.Println("Table: ", pt)

	default:
		fmt.Println("Please Enter Valid Number")
	}

}
