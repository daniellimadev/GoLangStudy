package main

import "fmt"

func main() {
	a := "e"
	b := "é"
	c := "U"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n\n", d, e, f)

	x := 20
	y := 20.0

	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T", y, y)

}
