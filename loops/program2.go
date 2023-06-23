// find the sum of the first 10 natural numbers
package main

import "fmt"

func main() {
	var n int = 10
	var sum1 int = 0

	for i :=1; i<=n; i++ {
		sum1 = sum1 + i
	}
	fmt.Print(sum1)
}