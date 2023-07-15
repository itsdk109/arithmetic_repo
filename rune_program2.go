// how to convert string to rune

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// convert string to rune
	str := "Mr. Deepak Kumar Choudhary"

	rune_str := []rune(str)
	fmt.Printf("Rune Value of this [%s] is :\n", str, rune_str)
	fmt.Println(" ")

	// count length of this string
	fmt.Println(len(str))
	fmt.Println("Count length with utf8.RuneCountInString -->", utf8.RuneCountInString(str))
}
