package main

import "fmt"

func main() {

	x := 9

	if !(x%2 == 0) && x%3 == 0 {
		fmt.Println("is a multiple of two and also of three")
	}

	if x%2 == 0 || x%3 == 0 {
		fmt.Println("is a multiple of two or three")
	}

}
