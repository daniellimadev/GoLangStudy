package main

import "fmt"

type client struct {
	Name    string
	Surname string
	Smoker  bool
}

type person struct {
	Name string
	Age  int
}

type professional struct {
	person
	Title  string
	Salary int
}

func main() {
	client1 := client{
		Name:    "João",
		Surname: "da Silva",
		Smoker:  false,
	}

	client2 := client{"Joana", "Pereira", true}

	fmt.Println(client1)
	fmt.Println(client2)

	person1 := person{
		Name: "Alfredo",
		Age:  30,
	}

	person2 := professional{
		person: person{
			Name: "Davi",
			Age:  31,
		},
		Title:  "pizza maker",
		Salary: 10000,
	}

	person3 := person{"Daniel", 20}
	person4 := professional{person{"Denis", 40}, "paid traffic manager", 100000}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3.Name)
	fmt.Println(person4.Name)

	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50}

	fmt.Println(x)
}
