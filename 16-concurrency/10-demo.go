package main

import (
	"fmt"
	"time"
)

func main() {
	ch := generateNumbers()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func generateNumbers() <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- 10
		time.Sleep(500 * time.Millisecond)
		ch <- 20
		time.Sleep(500 * time.Millisecond)
		ch <- 30
		time.Sleep(500 * time.Millisecond)
		ch <- 40
		time.Sleep(500 * time.Millisecond)
		ch <- 50
	}()
	return ch
}
