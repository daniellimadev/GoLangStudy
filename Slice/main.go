package main

import "fmt"

func main() {
	// array := [5]int{1, 2, 3, 4, 5}
	// fmt.Println(array)
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(slice)

	// slice2 := append(slice, 6)
	// fmt.Println(slice2)

	// fmt.Println(slice[3])
	// slice[3] = 348756
	// fmt.Println(slice[3])

	slice := []string{"banana", "apple", "jackfruit", "peach"}

	//for index, value := range slice {fmt.Println("At index", index, "we have the value:", value)}

	//slice[4] = "watermelon"
	slice = append(slice, "watermelon")

	for _, value := range slice {
		fmt.Printf("One of the values ​​of this slice is %v.\n", value)
	}
}
