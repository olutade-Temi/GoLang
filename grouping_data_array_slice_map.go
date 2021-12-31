package main

import "fmt"

func main() {
	//Array
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	for k, l := range x {
		fmt.Println(k, l)
	}

	//Slice composite literals
	a := []int{4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a[0]) //accessing by index position
	//using range to loop through a slice and access its index position

	for i, v := range a {
		fmt.Println(i, v)
	}

	//another way of printing the index and value as shown above using range

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	//slicing a slice
	fmt.Println(a[:])
	fmt.Println(a[1:])
	fmt.Println(a[1:3])

	//append to a slice
	b := []int{8, 9, 10, 11, 12}
	b = append(b, 77, 88, 99, 100)
	fmt.Println(b)

	y := []int{23, 24, 25, 26, 27}
	b = append(b, y...)
	fmt.Println(b)

	//delete from a slice
	b = append(b[:2], b[4:]...)
	fmt.Println(b)

	//using make to assign size and capacity to your slice
	g := make([]int, 10, 100)
	fmt.Println(g)
	fmt.Println(len(g))
	fmt.Println(cap(g))

	//multidimensional slice
	jb := []string{"james", "Bond", "Chocolate", "martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	xp := [][]string{jb, mp}
	fmt.Println(xp)

	//maps
	m := map[string]int{
		"James":    32,
		"Temitope": 22,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	//to check if a key exists

	if v, ok := m["Temitope"]; ok {
		fmt.Println("This is the result ", v)
	}

	//using range on a map and how to add to a map
	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	//delete from a map
	delete(m, "Temitope")
	fmt.Println(m)

	if v, ok := m["James"]; ok {
		fmt.Println("Value:", v)
		delete(m, "James")
	}
	fmt.Println(m)

}
