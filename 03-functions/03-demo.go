/* anonymous functions */
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	greetMsg := func(userName string) string {
		result := fmt.Sprintf("Hi %s, Have a nice day!\n", userName)
		return result
	}("Suresh")
	fmt.Print(greetMsg)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("100 + 200 =", result)

	quotient, remainder := func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
}
