package main

import (
	"encoding/json"
	"fmt"
)

type persn struct {
	First string
	Last  string
	Age   int
}

func main() {
	p3 := persn{
		First: "Temitope",
		Last:  "Olutade",
		Age:   22,
	}
	p4 := persn{
		First: "Taiwo",
		Last:  "Olutade",
		Age:   27,
	}

	people := []persn{p3, p4}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
