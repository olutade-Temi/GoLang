package main

import "fmt"

func main() {
	defer foot() // this runs after exiting all other functions due to the defer keyword
	bar("Oladapo")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
	j := variadic_param(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", j)

	//Anonymous function
	func(g int) {
		fmt.Println("The meaning of life", g)
	}(42)

	k := func(d int) {
		fmt.Println("I was born in the year", d)
	}
	k(1999)
}
func foot() {
	fmt.Println("Hello Temi")
}
func bar(s string) {
	fmt.Println("Hello", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, " says hello")
	b := false
	return a, b
}

func variadic_param(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding ", v, "to the total which is now", sum)
	}
	return sum

}
