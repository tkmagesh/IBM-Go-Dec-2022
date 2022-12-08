package main

import (
	"fmt"
)

//share memory by communicating

//Consumer
func main() {

	ch := add(100, 200)
	result := <-ch
	fmt.Println("Result :", result)
}

//Producer
func add(x, y int) chan int {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result
	}()
	return ch
}
