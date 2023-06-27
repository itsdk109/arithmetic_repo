// Multiple return values

package main
import "fmt"

func mul(x int , y int) (int,int) {
	var multiplication int = x * y
	var area_triangle int = (x*y)/2

	return area_triangle, multiplication
}

func main(){

a, m := mul(7,10) // calling Function
fmt.Println("Area of Triangle:",a," ", "Product of x*y: ", m)

}