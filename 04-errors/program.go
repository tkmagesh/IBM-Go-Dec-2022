package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	// fmt.Println(divide(100, 0))
	/*
		result, err := divide(100, 0)
		if err != nil {
			fmt.Println("something went wrong")
			fmt.Println(err)
			return
		}
		fmt.Println("Result :", result)
	*/
	result, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Please do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong")
		fmt.Println(err)
		return
	}
	fmt.Println("Result :", result)
}

/*
func divide(x, y int) (int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, err
	}
	result := x / y
	return result, nil
}
*/

/*
func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be 0")
		return
	}
	result = x / y
	return
}
*/

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
