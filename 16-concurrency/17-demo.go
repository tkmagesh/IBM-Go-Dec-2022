package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	fibCh := genFibonacci(done)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		// done <- struct{}{}
		close(done)
	}()

	for fibNo := range fibCh {
		fmt.Println(fibNo)
	}
}

func genFibonacci(done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
*/
