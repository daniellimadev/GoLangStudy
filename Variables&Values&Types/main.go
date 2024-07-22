package main

import "fmt"

//Primitive types
var x int
var y float64
var w bool

// Creating your own type
type daniel int

var d daniel = 10

func main() {
	x = 20
	fmt.Printf("%v, %T\n\n", x, x)
	x = 50
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n\n\n", w, w)

	fmt.Printf("%v\n", x)

	// converting one type to another
	x = int(d)
	fmt.Printf("%v\n", x)

}
