package main

import (
	"fmt"
)

func main() {
	var nos []int
	nos = []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an slice (using indexer)")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating an slice (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	fmt.Println("Adding new values to the slice")
	nos = append(nos, 10)
	nos = append(nos, 20, 30, 40)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}
	/*
		slicing
		nos[lo:hi] => from index lo to hi-1
		nos[lo] => from index lo to end of the slice
		nos[:hi] => from index 0 to hi - 1
	*/
	fmt.Println("nos[2:6] =", nos[2:6])
	fmt.Println("nos[:6] =", nos[:6])
	fmt.Println("nos[6:] =", nos[6:])

	// subList := nos[2:5] //creates a pointer to the underlying array
	// To create a copy, use the copy function
	subList := make([]int, 3)
	copy(subList, nos[2:5])
	fmt.Println("nos", nos)
	fmt.Println("subList", subList)
	subList[0] = 999
	fmt.Println("After modifying subList")
	fmt.Println("nos", nos)
	fmt.Println("subList", subList)
}
