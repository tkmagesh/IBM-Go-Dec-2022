/* higher order functions - functions as return values  (applied) */

package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

//Custom type
type OperationFn func(int, int)

func main() {
	composeProfile(composeLog(add))(100, 200)
}

func composeLog(operation OperationFn) OperationFn {
	return func(x, y int) {
		fnName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
		log.Printf("Operation started - %s(%d, %d)\n", fnName, x, y)
		operation(x, y)
		log.Println("Operation completed")
	}
}

func composeProfile(operation OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

// 3rd party library api
func add(x, y int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	time.Sleep(4 * time.Second)
	fmt.Println("Subtract Result :", x-y)
}
