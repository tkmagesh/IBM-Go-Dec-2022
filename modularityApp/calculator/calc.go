package calculator

import "fmt"

var op_count int

var ProductRanks map[string]int

func init() {
	fmt.Println("calculator-init invoked")
	ProductRanks = make(map[string]int)
}

func OpCount() int {
	return op_count
}
