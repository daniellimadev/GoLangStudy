package main

import "fmt"

func main() {

	friends := map[string]int{
		"Daniel": 5551234,
		"Joana":  9996674,
	}

	fmt.Println(friends)
	fmt.Println(friends["Joana"])

	friends["gopher"] = 444444

	fmt.Println(friends)
	fmt.Println(friends["gopher"], "\n\n")

	// comma ok idiom
	if itWillBe, ok := friends["Ghost"]; !ok {
		fmt.Println("it does not have!")
	} else {
		fmt.Println(itWillBe)
	}

	fmt.Println()

	anything := map[int]string{
		123: "very cool",
		98:  "a little less cool",
		983: "this is cool",
		19:  "old enough to go to a party",
	}

	fmt.Println(anything)

	total := 0

	for key, _ := range anything {
		total += key
	}

	fmt.Println(total)
}
