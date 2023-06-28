// String array

package main
import "fmt"

func main(){

	// array
	var c = [...] string {"Patna","Purana Bhojpur","UP","Delhi","Gujrat"}
	fmt.Println("Array of c is: ",c)

	// copy of c is assigned to d
	d := c
	// change "Purana Bhojpur" to "Mumbai"
	d[1] = "Mumbai"

	fmt.Println("Array of D is: ",d)

}