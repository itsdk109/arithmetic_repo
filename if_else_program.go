package main
import "fmt"

func main(){

	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d",&num)

	if num%2==0 {
		fmt.Println("Even Number")
	} else
		{
		fmt.Println("Odd Number")
		}
}