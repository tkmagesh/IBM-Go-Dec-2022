/* Using multiple channels */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		data3 := <-ch3
		fmt.Println("ch3 :", data3)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	/*
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			fmt.Println("ch1 :", <-ch1)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			fmt.Println("ch2 :", <-ch2)
			wg.Done()
		}()
		wg.Wait()
	*/
	for i := 0; i < 3; i++ {
		select {
		case data1 := <-ch1:
			fmt.Println("ch1 :", data1)
		case ch3 <- 300:
			fmt.Println("data sent to ch3")
		case data2 := <-ch2:
			fmt.Println("ch2 :", data2)
			/* default:
			fmt.Println("No channel operations succeeded") */
		}
	}

}
