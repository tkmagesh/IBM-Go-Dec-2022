/* function basics */
package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	fmt.Print(getGreetMsg("Suresh"))
	fmt.Println("Sum of 100 and 200 is ", add(100, 200))
	// fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	//attempting to print only the quotient
	/*
		quotient, remainder := divide(100, 7) // unable to compile coz the "remainder" is not used
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
	*/

	/*
		quotient := divide(100, 7) // assignment mismatch: 1 variable but divide returns 2 values
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
	*/

	quotient, _ := divide(100, 7) // "_" is used in place of any variable that we dont want to use
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
}

/* function with no arguments & no return values */
func sayHi() {
	fmt.Println("Hi there!")
}

/* function with 1 argument */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* function with 1 argument and 1 return result */
func getGreetMsg(userName string) string {
	result := fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
	return result
}

/* function with 2 arguments and 1 return result */
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/*
func add(x, y int) (result int) {
	result = x + y
	return
}
*/

/* function with more than 1 return result */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/* function with named return result */
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient, remainder = x/y, x%y
	return
}
