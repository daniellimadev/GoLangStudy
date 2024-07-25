package main

import "fmt"

func main() {
	for hours := 0; hours <= 12; hours++ {
		fmt.Println("Hour: ", hours)

		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Print("\n\n")
	}
}
