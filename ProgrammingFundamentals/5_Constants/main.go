package main

import "fmt"

const x = 10

var y = 10

var a int
var b float64

const (
	g = 16
	h = 45
	i = 67
)

func main() {
	a = x
	fmt.Println(a)
	fmt.Println(g, h, i)
}
