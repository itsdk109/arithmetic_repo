/* I'll  cover :
1. How to Create Map?
2. How to add item to a Map?
3. Retrieving a value of a Key from a Map
4. Check if a Key exists in a Map?.
5. Iterate Over all elements of a Map.
6. Deleting items from a map.
7. Check if Two maps are equall or Not?

*/

package main

import (
	"fmt"
)

func main() {

	/* How to create map */
	var salary = make(map[string]int)
	fmt.Println(salary) // Output: map[]

	/* or Second method*/

	var salary1 = map[string]int{}
	fmt.Println(salary1) // Output: map []

	/* ----------------------------------------------------------------- */

	/* How to Add item to a map */
	var salary2 = make(map[string]int)

	salary2["Deepak"] = 1000
	salary2["Rohit"] = 2000
	salary2["Ajay"] = 3000

	fmt.Println(salary2) // Output: map["Ajay":3000 "Deepak":1000 "Rohit":2000]

	/* or Second method*/
	var salary3 = map[string]int{
		"Deepak Kumar": 10,
		"Rohit Kumar":  20,
		"Ajay Kumar":   30,
	}
	fmt.Println(salary3) // Output: map["Ajay Kumar":30 "Deepak Kumar":10 "Rohit Kumar":20]

	/* ------------------------------------------------------------------- */

	/* Check if the key is exists or not */
	var salary4 = map[string]int{
		"A": 5,
		"B": 10,
		"C": 15,
	}

	keyTofind := "D"

	value, Key := salary4[keyTofind]
	if Key == true {
		fmt.Println(keyTofind, "exists in a map and the value is ", value)
		return
	}
	fmt.Println(keyTofind, "Does not exists in a map and value is ", value)
	/* ------------------------------------------------------------- */
	fmt.Println(" ")

	/* How to Iterate over all elememts of a map */

	var salary5 = map[string]int{
		"Xa": 5,
		"Ya": 10,
		"Za": 15,
	}
	fmt.Print("Get all Elements:\n")
	for name_key, data := range salary5 {
		fmt.Println("\t", name_key, data)
	}

	// delete
	deleted := "Xa"
	delete(salary5, deleted)
	fmt.Println(deleted, "is Deleted\n Final map is ", salary5)

	/* -------------------------------------------------------- */

}
