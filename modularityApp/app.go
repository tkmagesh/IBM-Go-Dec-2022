package main

import (
	"fmt"
	// "github.com/tkmagesh/ibm-go-dec-2022/modularityApp/calculator"
	// calc "github.com/tkmagesh/ibm-go-dec-2022/modularityApp/calculator"

	"github.com/fatih/color"
	_ "github.com/tkmagesh/ibm-go-dec-2022/modularityApp/calculator" // imported ONLY to execute the init() functions
	utils "github.com/tkmagesh/ibm-go-dec-2022/modularityApp/calculator/utils"
)

func main() {
	// fmt.Println("modularity-app invoked")
	color.Red("modularity-app invoked")
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("Op Count :", calculator.OpCount())
	*/
	/*
		fmt.Println(calc.Add(100, 200))
		fmt.Println(calc.Subtract(100, 200))
		fmt.Println("Op Count :", calc.OpCount()) */
	fmt.Println(utils.IsEven(21))
}
