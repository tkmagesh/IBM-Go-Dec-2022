package main

import "fmt"

func main() {
	const pi float64 = 3.14
	fmt.Println(pi)

	//iota
	/*
		const red int = 1
		const green int = 2
		const blue int = 3
	*/

	/*
		const (
			red   int = 1
			green int = 2
			blue  int = 3
		)
	*/

	/*
		const (
			red   int = iota
			green int = iota
			blue  int = iota
		)
	*/

	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red   = iota + 1 // 0 + 1
			green            // 1 + 1
			blue             // 2 + 1
		)
	*/

	/*
		const (
			red   = iota * 2 // 0 * 2
			green            // 1 * 2
			blue             // 2 * 2
		)
	*/

	const (
		red   = iota * 2 // 0 * 2
		green            // 1 * 2
		_                // 2 * 2
		blue             // 3 * 2
	)

	fmt.Printf("red = %d, green = %d, blue = %d\n", red, green, blue)

	//IOTA Usage:
	const (
		VERBOSE = 1 << iota
		CONFIG_FROM_DISK
		DATABASE_REQUIRED
		LOGGER_ACTIVATED
		DEBUG
		FLOAT_SUPPORT
		RECOVERY_MODE
		REBOOT_ON_FAILURE
	)
	//fmt.Printf("%b, %b, %b, %b, %b, %b, %b, %b\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
	fmt.Printf("%d, %d, %d, %d, %d, %d, %d, %d\n", VERBOSE, CONFIG_FROM_DISK, DATABASE_REQUIRED, LOGGER_ACTIVATED, DEBUG, FLOAT_SUPPORT, RECOVERY_MODE, REBOOT_ON_FAILURE)
}
