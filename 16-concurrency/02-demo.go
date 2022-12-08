package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the counter by 1
	go f1()   //schedule the function to be executed in future by the scheduler

	wg.Add(1)
	go f1()
	f2()
	wg.Wait() // block until the counter becomes 0
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrementing the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
