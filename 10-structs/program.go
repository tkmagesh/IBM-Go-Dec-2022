package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	// var emp Employee
	// var emp Employee = Employee{100, "Magesh", 10000} // do not use this
	// var emp Employee = Employee{Id: 100, Name: "Magesh", Salary: 10000}
	/*
		var emp Employee = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/
	/*
		var emp = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Printf("%+v\n", emp)

	// accessing the fields of the struct
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.Id, emp.Name, emp.Salary)

	/*
		var emp2 Employee = emp // creating a copy of the emp
		emp2.Salary = 20000
		fmt.Printf("emp : %+v\n", emp)
		fmt.Printf("emp2 : %+v\n", emp2)
	*/

	var emp2 *Employee = &emp
	// fmt.Printf("emp2 : Id = %d, Name = %q, Salary = %.2f\n", (*emp2).Id, (*emp2).Name, (*emp2).Salary) // No need to dereference the pointer to access the fields
	fmt.Printf("emp2 : Id = %d, Name = %q, Salary = %.2f\n", emp2.Id, emp2.Name, emp2.Salary)

	awardBonus(&emp, 10000)
	fmt.Println("After awarding 10000 bonus")
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.Id, emp.Name, emp.Salary)

}

func awardBonus(emp *Employee, bonus float32) {
	emp.Salary += bonus
}
