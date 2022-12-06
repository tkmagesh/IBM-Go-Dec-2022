package main

import "fmt"

func main() {
	var x int = 100
	var varPtr *int = &x
	fmt.Println(varPtr)

	//dereferencing (accessing the underlying value of a pointer)
	fmt.Println(*varPtr)

	fmt.Println(x == *(&x))

	fmt.Println("Before incrementing : x =", x)
	increment(&x)
	fmt.Println("After incrementing : x =", x)

	no1, no2 := 100, 200
	fmt.Printf("Before swapping, no1 = %d and no2 = %d\n", no1, no2)
	swap( /*  */ )
	fmt.Printf("After swapping, no1 = %d and no2 = %d\n", no1, no2)
}

func increment(no *int) {
	fmt.Println("&no =", no)
	*no++
}

func swap( /*  */ ) {
	/*  */
}
