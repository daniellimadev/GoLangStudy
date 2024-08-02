package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) goodmorning() {
	fmt.Println(p.Name, "say good morning!")
}

func main() {

	daniel := Person{"Daniel", 24}
	daniel.goodmorning()

}

// func (receiver) identifier(parameters) (returns) { code }
