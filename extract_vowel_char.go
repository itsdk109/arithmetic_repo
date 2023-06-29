// How to extract Vowel Char from given string.

package main

import "fmt"

func main() {

	// Extracting vowel character from given string
	var g_string string = "thequickbrownfoxjumpsoverthelazydog"

	fmt.Println("Given String: %s", g_string)

	for _, i := range g_string {
		if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' {
			fmt.Printf("%c is vowel\n", i)
		} else {
			fmt.Print(" ")
		}
	}
}
