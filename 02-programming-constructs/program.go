/* programming constructs */
package main

import "fmt"

func main() {
	/* if else */
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 { // the scope the variable is limited to the if-else block
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	/* switch case */
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10 => "Excellent"
		otherwise => "Invalid Score"
	*/
	/*
		score := 6
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")
		}
	*/

	/*
		score := 6
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not Bad")
		case 8, 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid Score")
		}
	*/

	switch score := 6; score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid Score")
	}

	// Using switch case like a "nested if"
	switch no := 22; {
	case no%2 == 0:
		fmt.Printf("%d is an even number\n", no)
	case no%2 != 0:
		fmt.Printf("%d is an odd number\n", no)
	}

	//fallthrough
	switch x := 4; x {
	case 0:
		fmt.Println("x is 0")
	case 1:
		fmt.Println("x <= 1")
		fallthrough
	case 2:
		fmt.Println("x <= 2")
		fallthrough
	case 3:
		fmt.Println("x <= 3")
		fallthrough
	case 4:
		fmt.Println("x <= 4")
		fallthrough
	case 5:
		fmt.Println("x <= 5")
		fallthrough
	case 6:
		fmt.Println("x <= 6")
		// fallthrough
	case 7:
		fmt.Println("x <= 7")
		fallthrough
	case 8:
		fmt.Println("x <= 8")
	}

	//fallthrough - applied
	switch plan := "Super"; plan {
	case "Super":
		fmt.Println("All Super features")
		fallthrough
	case "Premium":
		fmt.Println("All Premium features")
		fallthrough
	case "Pro":
		fmt.Println("All Pro features")
		fallthrough
	case "Free":
		fmt.Println("All Free features")

	}

	/* for */
	/* ver 1.0 */
	fmt.Println("for (v1.0)")
	for idx := 0; idx < 10; idx++ {
		fmt.Println(idx)
	}

	/* ver 2.0 (while)*/
	fmt.Println("for (v2.0) [like while]")
	/*
		no := 1
		for no < 100 {
			no += no
		}
		fmt.Printf("no = %d\n", no)
	*/

	for no := 1; no < 100; {
		no += no
		fmt.Printf("no = %d\n", no)
	}
	// fmt.Printf("no = %d\n", no) // "no" variable is not accessible outside the "for" block

	fmt.Println("for (v3.0) [infinite]")
	x := 1
	/*
		for {
			x += x
			if x > 100 {
				break
			}
		}
	*/
	for {
		x += x
		if x < 100 {
			continue
		}
		break
	}

	fmt.Printf("x = %d\n", x)

	fmt.Println("Using Labels")
OUTER_LOOP:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("=======================")
				continue OUTER_LOOP
			}
		}
	}
}
