package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schedule the function to be executed in future by the scheduler
	f2()
	// time.Sleep(100 * time.Millisecond) // DO NOT do this
	fmt.Scanln() // DO NOT do this
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
