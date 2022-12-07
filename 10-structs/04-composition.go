package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product
	// Dummy
	Expiry string
	// Name   string // overriding Product.Name
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{Id: id, Name: name, Cost: cost},
		Expiry:  expiry,
	}
}

func main() {
	// grapes := PerishableProduct{Product{Id: 100, Name: "Grapes", Cost: 50}, "2 Days"}
	/*
		grapes := PerishableProduct{
			Product: Product{Id: 100, Name: "Grapes", Cost: 50},
			Expiry:  "2 Days",
		}
	*/
	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")
	// fmt.Println(grapes.Product.Id, grapes.Product.Name, grapes.Product.Cost, grapes.Expiry)
	fmt.Println(grapes.Id, grapes.Name, grapes.Cost, grapes.Expiry)

	//pointer of an empty object
	//pPtr := &PerishableProduct{}
	pPtr := new(PerishableProduct) // equivalent to &PerishableProduct{}
	fmt.Println(pPtr)

}
