package main

import "fmt"

func main() {

	// Memory -> Address -> Value
	a := 150
	var pointer *int = &a

	*pointer = 50
	b := &a
	*b = 60

	fmt.Println(a)

	// When to use pointers

	x := 25

	// isreceivingvalue(x)

	isreceivingpointer(&x)
}

func isreceivingvalue(x int) {
	x++
	fmt.Println("In function:", x)

}

func isreceivingpointer(x *int) {
	*x++
	fmt.Println("In function:", *x)
}
