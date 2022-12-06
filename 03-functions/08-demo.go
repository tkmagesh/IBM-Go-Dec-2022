/* higher order functions - functions as return values  (applied) */

package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

func main() {
	/*
		log.Println("Operation started")
		add(100, 200)
		log.Println("Operation completed")

		log.Println("Operation started")
		subtract(100, 200)
		log.Println("Operation completed")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	// Only logging
	/*
		logAdd := composeLog(add)
		logSubtract := composeLog(subtract)
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		composeLog(add)(100, 200)
		composeLog(subtract)(100, 200)
	*/

	// Only profiling
	/*
		profileAdd := composeProfile(add)
		profileAdd(100, 200)
	*/

	/*
		logAdd := composeLog(add)
		profileLogAdd := composeProfile(logAdd)
		profileLogAdd(100, 200)
	*/

	composeProfile(composeLog(add))(100, 200)

}

func composeLog(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		fnName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
		log.Printf("Operation started - %s(%d, %d)\n", fnName, x, y)
		operation(x, y)
		log.Println("Operation completed")
	}
}

func composeProfile(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

/*
func logOperation(x, y int, operation func(int, int)) {
	fnName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
	log.Printf("Operation started - %s(%d, %d)\n", fnName, x, y)
	operation(x, y)
	log.Println("Operation completed")
}
*/

/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}
*/

// 3rd party library api
func add(x, y int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	time.Sleep(4 * time.Second)
	fmt.Println("Subtract Result :", x-y)
}
