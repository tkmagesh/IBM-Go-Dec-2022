package main

import (
	"fmt"
	"sort"
)

/*
	Write the apis for the following

	IndexOf => return the index of the given product
		ex:  returns the index of the given product

	Includes => return true if the given product is present in the collection else return false
		ex:  returns true if the given product is present in the collection

	Filter => return a new collection of products that satisfy the given condition
		use cases:
			1. filter all costly products (cost > 1000)
				OR
			2. filter all stationary products (category = "Stationary")
				OR
			3. filter all costly stationary products
			etc

	Any => return true if any of the product in the collections satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			2. are there any stationary product (category = "Stationary")?
			3. are there any costly stationary product?
			etc

	All => return true if all the products in the collections satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			2. are all the products stationary products (category = "Stationary")?
			3. are all the products costly stationary products?
			etc


	Sort => Sort the products collection by any attribute
		IMPORTANT : MUST Use sort.Sort() function
		use cases:
			1. sort the products collection by cost
			2. sort the products collection by name
			3. sort the products collection by units
			etc

*/
type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) Print(label string) {
	fmt.Println(label)
	for _, product := range products {
		fmt.Println(product)
	}
}

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

//sort.Interface implementation
func (products Products) Len() int {
	return len(products)
}

//to sort by ID
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

// to sort by Name
type ByName struct {
	Products
}

//overriding the implementation of the base type (Products)
func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

// to sort by Cost
type ByCost struct {
	Products
}

//overriding the implementation of the base type (Products)
func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

// to sort by Name
type ByUnits struct {
	Products
}

//overriding the implementation of the base type (Products)
func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

// to sort by Name
type ByCategory struct {
	Products
}

//overriding the implementation of the base type (Products)
func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

// helper method for sorting
func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	}
}

//using the sort.Slice() function
func (products Products) Sort2(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	case "Name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	products.Print("Initial List")
	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println("Index of Stove :", products.IndexOf(stove))

	fmt.Println("Filter")
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	costlyProducts.Print("Costly Products")

	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	stationaryProducts.Print("Stationary Products")

	costlyStationaryProductPredicate := func(product Product) bool {
		return costlyProductPredicate(product) && stationaryProductPredicate(product)
	}
	costlyStationaryProducts := products.Filter(costlyStationaryProductPredicate)
	costlyStationaryProducts.Print("Costly & Stationary Products")

	fmt.Println("Sorting")
	/*
		sort.Sort(products)
		products.Print("Default Sort (by Id)")

		sort.Sort(ByName{products})
		products.Print("By Name")

		sort.Sort(ByCost{products})
		products.Print("By Cost")

		sort.Sort(ByUnits{products})
		products.Print("By Units")

		sort.Sort(ByCategory{products})
		products.Print("By Category")
	*/
	// using the "Sort" helper method
	// products.Sort("Id")
	products.Sort2("Id")
	products.Print("Default Sort (by Id)")

	// products.Sort("Name")
	products.Sort2("Name")
	products.Print("By Name")

	// products.Sort("Cost")
	products.Sort2("Cost")
	products.Print("By Cost")

	// products.Sort("Units")
	products.Sort2("Units")
	products.Print("By Units")

	// products.Sort("Category")
	products.Sort2("Category")
	products.Print("By Category")
}
