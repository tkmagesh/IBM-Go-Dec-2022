package main

import "fmt"

//type alias
type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Eiusmod amet aute ut ad aute officia nostrud. Ipsum labore ea laboris laboris mollit. Incididunt eu enim excepteur pariatur sit irure quis ex eiusmod. Duis dolor esse tempor voluptate eiusmod. Amet proident occaecat occaecat velit irure est fugiat id do amet quis minim mollit. Aliqua cupidatat ipsum excepteur ut adipisicing anim sint adipisicing culpa veniam nisi consequat.")
	fmt.Println(str.Length())

	//type conversion
	var i int32 = 100
	var f float64 = float64(i)
	fmt.Println(f)
}
