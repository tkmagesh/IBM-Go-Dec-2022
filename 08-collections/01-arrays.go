package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	// nos := [5]int{}
	fmt.Println(nos)

	fmt.Println("Iterating an array (using indexer)")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] = %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating an array (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}
}
