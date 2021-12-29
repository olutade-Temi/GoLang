package main

import "fmt"

func main() {
	s := "Temitope is a boy"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	sb := []byte(s)
	fmt.Println(sb)
	fmt.Printf("%T\n", sb)

}
