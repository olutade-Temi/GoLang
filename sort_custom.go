package main

import (
	"fmt"
	"sort"
)

type perso struct {
	first string
	age   int
}
type ByAge []perso

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	p1 := perso{"James", 32}
	p2 := perso{"Moneypenny", 27}
	p3 := perso{"Q", 64}
	p4 := perso{"M", 56}

	people := []perso{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
