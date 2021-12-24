package main

import "fmt"

func main() {
	fmt.Println("Hello World, this is my first interaction with GO Programming Language")
	foo()

	for i := 0
		i < 100;
	i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("This is just another line")
}
