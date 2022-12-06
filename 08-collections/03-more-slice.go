package main

import "fmt"

func main() {
	// var nos []int
	// nos := make([]int, 0, 5)
	nos := make([]int, 2, 5)
	//nos[0] = 100
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 10)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 20)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 30)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 40)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 50)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 60)
	fmt.Printf("len = %d, capacity = %d, nos = %v\n", len(nos), cap(nos), nos)

}
