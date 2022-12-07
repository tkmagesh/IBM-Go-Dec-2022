package main

import "fmt"

type Product struct {
}

//fmt.Stringer implementation
func (p Product) String() string {
	return "I am a product"
}

func main() {
	// var x interface{} //before go 1.18
	var x any // from go 1.18
	x = 100
	x = "Eiusmod et velit proident veniam tempor reprehenderit Lorem elit sunt ullamco nostrud in ea id. Eiusmod non non est Lorem ex irure veniam aliqua in ex eu nostrud eu cillum. Ullamco aliquip qui dolore sint consectetur duis magna consequat magna enim. Commodo aliquip exercitation sint veniam culpa eu elit laboris minim pariatur duis non."
	x = 99.98
	x = true
	x = struct{}{}
	fmt.Println("x =", x)

	// x = true
	x = 1000
	// y := x.(int) + 2000
	if val, ok := x.(int); ok {
		y := val + 2000
		fmt.Println(y)
	} else {
		fmt.Println("x is NOT an int")
	}

	// x = "Ut quis consequat irure aliquip veniam reprehenderit tempor labore id laboris pariatur. Est dolore anim laborum quis reprehenderit tempor irure aute aliquip. Deserunt proident laborum commodo consectetur veniam commodo tempor sunt proident laborum officia esse qui. Voluptate sit dolore labore eu magna esse officia deserunt. Enim labore incididunt ipsum labore labore elit sunt dolore in. Dolor veniam laborum eu qui enim."
	// x = true
	x = Product{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int. x + 2000 =", val+2000)
	case string:
		fmt.Println("x is a string. len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64, x * 10% =", val*0.1)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is an empty struct{}")
	case fmt.Stringer:
		fmt.Println("x is fmt.Stringer implementation")
	default:
		fmt.Println("Unknown type")
	}
}
