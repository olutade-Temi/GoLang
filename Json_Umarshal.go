package main

import (
	"encoding/json"
	"fmt"
)

type personal struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{
		"First":"Temitope", "Last":"Olutade", "Age":22
	}, {
		"First":"Taiwo", "Last":"Olutade", "Age":27
	}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}
	var people []personal

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nAll of the data", people)

	for i, v := range people {
		fmt.Println("\n PERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
