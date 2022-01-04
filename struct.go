package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "Oladapo",
			last:  "Olutade",
			age:   32,
		},
		ltk: true,
	}
	p1 := person{
		first: "Temitope",
		last:  "Olutade",
		age:   23,
	}

	//anonymous struct
	p2 := struct {
		name  string
		class string
		sex   string
	}{
		name:  "Temitope Olutade",
		class: "Graduate",
		sex:   "Male",
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p2)
	fmt.Println(sa1)
}
