package main

import "fmt"

func uniqueslice(a, b []string) []string {
	var result []string
	checkifstringpresent := make(map[string]bool)

	for _, name := range a {
		if !checkifstringpresent[name] {
			checkifstringpresent[name] = true
			result = append(result, name)
		}
	}

	for _, name := range b {
		if !checkifstringpresent[name] {
			checkifstringpresent[name] = true
			result = append(result, name)
		}
	}

	return result
}

func main() {
	fmt.Println("Merged Slice: ", uniqueslice([]string{"Deepak", "Ajay", "Akshay"}, []string{"Ajay", "Anuj", "Soni"}))
}
