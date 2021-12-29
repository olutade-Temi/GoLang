package main

import "fmt"

func main() {
	s := "Temitope is a boy"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	c := "H"
	fmt.Println(c)
	
	cs := []byte(c)
	fmt.Println(cs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
}
