package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 4, "pencil": 1, "marker": 2}
	/*
		var productRanks map[string]int = map[string]int{
			"pen":    4,
			"pencil": 1,
			"marker": 2,
		}
	*/
	/*
		productRanks := map[string]int{
			"pen":    4,
			"pencil": 1,
			"marker": 2,
		}
	*/
	/*
		productRanks := make(map[string]int)
		productRanks["pen"] = 4
	*/
	productRanks := map[string]int{
		"pen":    4,
		"pencil": 1,
		"marker": 2,
	}
	fmt.Println(productRanks)

	fmt.Println("# of products :", len(productRanks))

	fmt.Println("Iterating a map (using range)")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Check for the existence of a key")
	keyToCheck := "notebook"
	if val, ok := productRanks[keyToCheck]; ok {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, val)
	} else {
		fmt.Printf("key %q does not exist\n", keyToCheck)
	}

	fmt.Println("Deleting a key")
	keyToDelete := "notebook"
	delete(productRanks, keyToDelete)
	fmt.Println("After deleting ", keyToDelete)
	fmt.Println(productRanks)

	addProductRank("stapler", 10, productRanks)
	fmt.Println(productRanks)
}

func addProductRank(product string, rank int, productRanks map[string]int) {
	productRanks[product] = rank
}
