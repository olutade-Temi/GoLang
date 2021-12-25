package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	s := fmt.Sprintf("%v\t%v\t%v\t", a, b, c)
	fmt.Println(s)
}
