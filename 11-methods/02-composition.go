package main

/*
	Create a "ApplyDiscount" method for the "Product" which will update the product cost after applying the given discount %

	apply 10% discount on grapes (PerishableProduct) and print the grapes object

*/
import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", p.Id, p.Name, p.Cost)
}

type PerishableProduct struct {
	Product
	Expiry string
}

// overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	// return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Expiry = %q", pp.Id, pp.Name, pp.Cost, pp.Expiry)
	//reusing the Product.Format()
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
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
	fmt.Println(pen.Format())

	grapes := NewPerishableProduct(101, "Grapes", 50, "2 Days")
	fmt.Println(grapes.Format())
}
