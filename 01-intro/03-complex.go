package main

import "fmt"

func main() {
	var c1 complex128
	c1 = 4 + 5i
	fmt.Println(c1)
	fmt.Println("real =", real(c1))
	fmt.Println("imaginary =", imag(c1))

	c2 := 7 + 9i
	result := c1 + c2
	fmt.Println(result)
}
