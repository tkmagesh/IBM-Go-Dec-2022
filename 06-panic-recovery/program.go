package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic raised", err)
			fmt.Println("===========================")
			fmt.Println("Call stack")
			fmt.Println("===========================")
			debug.PrintStack()
			fmt.Println("===========================")
		}
		fmt.Println("Thank you!")
	}()
	result, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("Error in divide operation. ", err)
		return
	}
	fmt.Println("Result : ", result)

}

func divideClient(x, y int) (result int, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = er.(error)
			return
		}
	}()
	result = divide(x, y)
	return
}

//3rd party library
func divide(x, y int) int {
	if y == 0 {
		panic(errors.New("divisor cannot be 0"))
	}
	return x / y
}
