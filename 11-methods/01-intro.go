package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func (e Employee) WhoAmI() {
	fmt.Println("I am ", e.Name)
}

// As a function
/*
func AwardBonus(emp *Employee, bonus float32) {
	emp.Salary += bonus
}
*/

// As a method
func (emp *Employee) AwardBonus(bonus float32) {
	emp.Salary += bonus
}

func main() {
	e := Employee{Name: "Magesh", Salary: 10000}
	e.WhoAmI()
	fmt.Println(e)
	// AwardBonus(&e, 5000)
	// (&e).AwardBonus(5000)
	e.AwardBonus(5000)
	fmt.Println(e)
}
