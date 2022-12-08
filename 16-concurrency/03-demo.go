package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(7)
	wg := &sync.WaitGroup{}
	noOfGoroutines := flag.Int("count", 0, "Number of goroutines to execute")
	flag.Parse()

	fmt.Println("Hit ENTER to start")
	fmt.Scanln()
	for i := 1; i <= *noOfGoroutines; i++ {
		wg.Add(1)    // increment the counter by 1
		go f1(i, wg) //schedule the function to be executed in future by the scheduler
	}

	f2()
	wg.Wait() // block until the counter becomes 0
}

func f1(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the counter by 1
	fmt.Printf("f1[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("f1[%d] completed\n", id)
	return
}

func f2() {
	fmt.Println("f2 invoked")
}
