package main

import "fmt"

func main() {
	x := 0

	for x <= 20 {
		x++
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}

	fmt.Println()

	f := 0

	for {
		if f < 10 {
			f++
			fmt.Println("F is less than 10")
		} else {
			fmt.Println("F is greater than 10")
			break
		}
	}

	fmt.Println("The loop finished!")
}
