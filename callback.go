package main

import "fmt"

func main() {
	ii := []int{1, 3, 4, 2, 5, 6, 7, 8, 9}
	s := newf(ii...)
	fmt.Println(s)

	s2 := even(newf, ii...)
	fmt.Println("Even Numbers", s2)

	s3 := odd(newf, ii...)
	fmt.Println("Odd Numbers", s3)
}

func newf(xj ...int) int {
	totaal := 0
	for _, v := range xj {
		totaal += v
	}
	return totaal
}
func even(f func(xj ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)

}

func odd(f func(xj ...int) int, vi ...int) int {
	var y2 []int
	for _, v := range vi {
		if v%2 != 0 {
			y2 = append(y2, v)
		}
	}
	return f(y2...)
}
