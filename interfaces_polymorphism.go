package main

import (
	"fmt"
)

type personn struct {
	first string
	last  string
}

type secretAgentt struct {
	personn
	ltk bool
}

func (s secretAgentt) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgentt speak")
}

func (p personn) speak() {
	fmt.Println("I am", p.first, p.last, " - the personn speak")
}

type human interface {
	speak()
}

//Assertion
func barr(h human) {
	switch h.(type) {
	case personn:
		fmt.Println("I was passed into barrrrrr", h.(personn).first)
	case secretAgentt:
		fmt.Println("I was passed into barrrrrr", h.(secretAgentt).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgentt{
		personn: personn{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgentt{
		personn: personn{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := personn{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	barr(sa1)
	barr(sa2)
	barr(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
