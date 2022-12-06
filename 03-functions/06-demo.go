/* higher order functions - pass functions as arguments (applied) */

package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
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

	logOperation(100, 200, add)
	logOperation(100, 200, subtract)
}

func logOperation(x, y int, operation func(int, int)) {
	log.Printf("Operation started - %s(%d, %d)\n", GetFunctionName(operation), x, y)
	operation(x, y)
	log.Println("Operation completed")
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

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
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
