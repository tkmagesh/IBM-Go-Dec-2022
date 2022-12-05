package main

import "fmt"

func main() {
	/*
		var name string
		fmt.Println("Enter your name :")
		fmt.Scanln(&name)
		fmt.Println(name)
	*/
	/*
		var firstName, lastName string
		fmt.Println("Enter your full name :")
		fmt.Scanln(&firstName, &lastName)
		fmt.Printf("%s %s\n", firstName, lastName)
	*/
	var n1, n2 int
	fmt.Println("Enter the numbers : ")
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1 + n2)
}
