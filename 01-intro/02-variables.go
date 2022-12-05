package main

import "fmt"

// var msg string

// var msg string = "Hello World!"

// type inference
// var msg = "Hello World!"

// msg := "Hello World!" //=> NOT ALLOWED

// unused variable
// var x int = 100

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/
	/*
		var msg string = "Hello World!"
	*/

	// type inference
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!"
	/*
		:= can be used ONLY with function scoped variables (not with package scoped variables)
	*/
	fmt.Println(msg)

	// trying to have unused variables
	/*
		x := 100
		x += 200 // x = x + 200
	*/

	// Handling multiple variables

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 and 200 is"
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of 100 and 200 is"
		var result int = x + y
		fmt.Println(str, result)
	*/

	// type inference
	/*
		var x, y = 100, 200
		var str = "Sum of 100 and 200 is"
		var result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of 100 and 200 is"
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of 100 and 200 is"
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		x, y, str := 100, 200, "Sum of 100 and 200 is"
		result := x + y
		fmt.Println(str, result)
	*/

	//using the formatting verbs with Printf function
	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

}
