package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was not 41")
	} else {
		fmt.Println("our value was not 40")
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//switch statement
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print ?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true")
		fallthrough
	default:
		fmt.Println("This is default")
	}

	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("Miss money or bond or do no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("This is default")
	}
}
