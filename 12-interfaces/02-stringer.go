package main

/* Using the fmt.Stringer interface */

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

/*
func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", p.Id, p.Name, p.Cost)
}
*/

//fmt.Stringer interface implementation
func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discount float32) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

/*
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}
*/

//fmt.Stringer interface implementation
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product, pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{Id: id, Name: name, Cost: cost},
		Expiry:  expiry,
	}
}

func main() {

	//Product
	pen := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(pen.Format())
	fmt.Println(pen) //Print function will call the String() method (fmt.Stringer interface)
	grapes := NewPerishableProduct(101, "Grapes", 50, "2 Days")
	// fmt.Println(grapes.Format())
	fmt.Println(grapes) //Print function will call the String() method (fmt.Stringer interface)
	fmt.Println("After applying 10% discount")
	grapes.ApplyDiscount(10)
	// fmt.Println(grapes.Format())
	fmt.Println(grapes) //Print function will call the String() method (fmt.Stringer interface)
}
