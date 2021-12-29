package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {

		fmt.Printf("The outer loop %d\n ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop %d\n", j)
		}
	}

	//using for in place of a while statement
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.\n")

	//using for with if statement
	d := 1
	for {

		if d > 9 {
			break
		}
		fmt.Println(d)
		d++
	}
	fmt.Println("done.\n")

	//using break and continue to print even numbers
	g := 1
	for {
		g++
		if g > 50 {
			break
		}
		if g%2 != 0 {
			continue
		}
		fmt.Println(g)
	}

	for k := 33; k <= 122; k++ {
		fmt.Printf("%v\t%#x\t%#U\n", k, k, k)
	}
}
