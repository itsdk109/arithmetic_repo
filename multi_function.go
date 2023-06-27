// Multiple return values

package main
import "fmt"

func mul(x int , y int) (int,int) {
	var multiplication int = x * y
	var area_triangle int = (x*y)/2

	return area_triangle, multiplication
}

func main(){

area_triangle, multiplication := mul(7,10)
fmt.Println("Area of Triangle:",area_triangle," ", "Product of x*y: ", multiplication)

}