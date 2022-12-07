package main

import "fmt"

func main() {
	product := struct {
		Id   int
		Name string
	}{100, "pen"}
	fmt.Printf("%#v\n", product)
	/*
		fmt.Println(product.Id)
		fmt.Println(product.Name)
	*/
	printProduct(product)

	emptyObj := struct{}{}
	fmt.Println(emptyObj)
}
func printProduct(p struct {
	Id   int
	Name string
}) {
	fmt.Printf("Product : Id = %d, Name = %q\n", p.Id, p.Name)
}
