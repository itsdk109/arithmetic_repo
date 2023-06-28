package main
import "fmt"

func main(){

var n int = 0
fmt.Print("enter number: ")
fmt.Scanf("%d", &n)
result := odd_even(n)
fmt.Println(result)
}

func odd_even(n int) int {
	var result int
	
	if n%2==0{
	fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
return result
}