package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sun2(5, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)

	fmt.Println(sun1(34, 65))
	fmt.Println(sun3(1, 2, 3, 4, 6, 77, 85, 9, 56))

	// Closures Function
	total := func() int {
		return sun3(1, 2, 3, 4, 6, 77, 85, 9, 56) * 4
	}()

	fmt.Println(total)
}

// Function with return. Closures
func sun1(a, b int) int {
	return a + b
}

// Function with errors
func sun2(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("The sum is greater than 50")
	}
	return a + b, nil
}

// Function with multiple returns.
func sun3(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
