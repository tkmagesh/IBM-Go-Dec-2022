package main

import (
	"fmt"
	"time"
)

func main() {
	fibCh := genFibonacci()
	for fibNo := range fibCh {
		fmt.Println(fibNo)
	}
}

func genFibonacci() <-chan int {
	ch := make(chan int)
	timeoutCh := timeout(10 * time.Second)

	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-timeoutCh:
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

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
