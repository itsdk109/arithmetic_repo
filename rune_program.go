// How to create rune with user define function

package main

import "fmt"

func rune_func(s rune) rune {
	return s
}

func main() {
	var s rune
	s = 'g'

	// calling function
	result1 := rune_func(s)
	fmt.Println(result1)
}
