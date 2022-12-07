package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    Organization
}

func main() {
	ibmLabs := Organization{
		Name: "IBM",
		City: "Cochin",
	}
	e1 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    ibmLabs,
	}
	fmt.Printf("e1 :=> %+v\n", e1)

	e2 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    ibmLabs,
	}
	fmt.Printf("e2 :=> %+v\n", e2)

	fmt.Println("After chaning Org City to Kollam")
	ibmLabs.City = "Kollam"
	fmt.Printf("e1 :=> %+v\n", e1)
	fmt.Printf("e2 :=> %+v\n", e2)

}
