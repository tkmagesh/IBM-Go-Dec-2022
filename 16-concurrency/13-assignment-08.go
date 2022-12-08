package main

import (
	"fmt"
	"time"
)

/*
Write a go routine that will generate prime numbers between the given start and end
Print the prime numbers in the main function as and when they are generated
*/

func main() {
	primeNoCh := generatePrimes(3, 100)
	for primeNo := range primeNoCh {
		fmt.Println(primeNo)
	}
}

func generatePrimes(start, end int) chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				ch <- no
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
