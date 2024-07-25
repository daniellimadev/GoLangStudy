package main

import "fmt"

func main() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("x is less than 5")
	case x == 5:
		fmt.Println("x is equal to 5")
	case x > 5:
		fmt.Println("x is greater than 5")
	}
}
