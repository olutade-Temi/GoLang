package main

import "fmt"

func main() {
	s1 := fool()
	fmt.Println(s1)

	xd := bart()
	fmt.Printf("%T\n", xd)
	fmt.Println(xd())
	//another way
	fmt.Println(bart()())
}

func fool() string {
	s := "Hello World"
	return s
}

func bart() func() int {
	return func() int {
		return 451
	}
}
