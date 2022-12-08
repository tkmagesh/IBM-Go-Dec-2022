package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	fmt.Println("before sending data, len(ch) :", len(ch))
	ch <- 100
	fmt.Println("after sending data, len(ch) :", len(ch))
	data := <-ch
	fmt.Println("after receiving data, len(ch) :", len(ch))
	fmt.Println(data)
}
