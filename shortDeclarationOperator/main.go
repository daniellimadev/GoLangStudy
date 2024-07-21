package main

import "fmt"

var z = "Hello, Go"

func main() {
	x := 10 + 30
	y := 25.
	r := 20 == 200

	fmt.Println(x)
	fmt.Println(r)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	x, w := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("w: %v, %T\n", w, w)
}
