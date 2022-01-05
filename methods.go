package main

import (
	"fmt"
)

type person1 struct {
	first string
	last  string
}

type secretAgent1 struct {
	person1
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }

func (s secretAgent1) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent1{
		person1: person1{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent1{
		person1: person1{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
