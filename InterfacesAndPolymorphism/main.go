package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Age     int
}

type Dentist struct {
	Person
	Teethpulled int
	Salary      float64
}

type Architect struct {
	Person
	Constructiontype string
	Sizeofmadness    string
}

func (x Dentist) goodmorning() {
	fmt.Println("My name is", x.Name, "and I have already pulled", x.Teethpulled, "teeth, and listen: Good morning!")
}

func (x Architect) goodmorning() {
	fmt.Println("My name is", x.Name, "and listen: Good morning!")
}

type People interface {
	goodmorning()
}

func humanbeing(g People) {
	g.goodmorning()
	switch g.(type) {
	case Dentist:
		fmt.Println("I earn:", g.(Dentist).Salary)

	case Architect:
		fmt.Println("I build:", g.(Architect).Constructiontype)
	}
}

func main() {
	mrtooth := Dentist{
		Person: Person{
			Name:    "Mister Tooth",
			Surname: "da Silva",
			Age:     50,
		},
		Teethpulled: 8973,
		Salary:      98736.06,
	}

	mrbuilding := Architect{
		Person: Person{
			Name:    "Mister Building",
			Surname: "Last Name",
			Age:     51,
		},
		Constructiontype: "Hospices",
		Sizeofmadness:    "Too much",
	}

	humanbeing(mrtooth)
	fmt.Println("")
	humanbeing(mrbuilding)

}
