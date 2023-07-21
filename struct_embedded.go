package main

import "fmt"

// create struct
type Employee struct {
	Name            string
	DestinationType string
	EmployeeType    string
}

// create another struct for embedding

type Company struct {
	Employee    // embedded struct
	CompanyName string
	Location    string
}

func main() {

	// create an instance of Comapany struct and access both Employee and Company fields/methods
	com_details := Company{
		Employee: Employee{
			Name:            "Deepak Kumar Choudhary",
			DestinationType: "Backend Developer in Go Language",
			EmployeeType:    "Internship",
		},
		CompanyName: "TechnoBuds IT Services and Consultancy Company",
		Location:    "Sector 105 Noida, Noida, IN",
	}

	fmt.Printf("%+v\n", com_details)

	fmt.Println(" ")

	// accessing specififc fiels of struct
	fmt.Printf("%+v\n", com_details.Name)
	fmt.Printf("%+v\n", com_details.DestinationType)
	fmt.Printf("%+v\n", com_details.CompanyName)

}
