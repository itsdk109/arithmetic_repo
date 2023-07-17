package main

import "fmt"

type mobile struct {
	BrandName string
	Processor string
	Price     int
}

func access_struct(brandname string, processor string, price int) mobile {
	return mobile{
		BrandName: brandname,
		Processor: processor,
		Price:     price,
	}
}

func main() {
	brand := access_struct("Oppo", "SnapDragon", 25486)
	fmt.Printf("%+v\n", brand)
}
