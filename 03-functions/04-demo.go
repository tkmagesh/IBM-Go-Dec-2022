/* higher order functions - assign functons as values to variables */

package main

import "fmt"

func main() {
	/*
		var sayHi func()
		sayHi = func() {
			fmt.Println("Hi there!")
		}
		sayHi()
	*/

	var fn func()
	fn = func() {
		fmt.Println("Hi there!")
	}
	fn()

	fn = func() {
		fmt.Println("Hello there!")
	}
	fn()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var operation func(int, int) int

	operation = func(i1, i2 int) int {
		return i1 + i2
	}
	fmt.Println(operation(100, 200))

	operation = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Println(operation(100, 200))

}
