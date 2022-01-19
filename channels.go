package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42

	}()

	fmt.Println("-----------")
	cr := make(<-chan int) //receive
	cd := make(chan<- int) //send

	fmt.Println(<-c)
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cd\t%T\n", cd)

}
