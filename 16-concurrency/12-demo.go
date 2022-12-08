package main

import (
	"fmt"
	"math/rand"
	"time"
)

//consumer
func main() {
	ch := generateNumbers()
	for val := range ch {
		fmt.Println(val)
	}
}

//producer
func generateNumbers() <-chan int {
	ch := make(chan int)
	max := rand.Intn(30)
	go func() {
		for i := 1; i <= max; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
