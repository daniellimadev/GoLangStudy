package main

import "fmt"

// Relational operators (==, <=, >=, !=, <, >)

var x bool

func main() {
	fmt.Println(x) // Zero value

	x = true
	fmt.Println(x) // Assigned value

	x := (20 == 100)
	fmt.Println(x) // Making the comparison

}
