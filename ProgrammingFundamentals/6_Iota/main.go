package main

import "fmt"

const (
	_ = iota * 5
	y
	z
)

func main() {
	fmt.Println(y, z)
}
