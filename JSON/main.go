package main

import (
	"encoding/json"
	"fmt"
)

// Defining a structure that matches JSON
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Unmarshal

	jsonData := `{"name": "Daniel Pereira", "age": 20, "email": "daniel@example.com"}`

	// Create an instance of the Person struct
	var person Person

	// Deserialize JSON to struct
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		panic(err)
	}

	// Displays the decoded data
	fmt.Printf("Nome: %s\nIdade: %d\nEmail: %s\n\n", person.Name, person.Age, person.Email)

	// Marshal

	// Instantiation of a struct
	person2 := Person{
		Name:  "Allan",
		Age:   25,
		Email: "allan@example.com",
	}

	// Serialize a struct to JSON
	jsonData2, err := json.Marshal(person2)
	if err != nil {
		panic(err)
	}

	// Display JSON as a string
	fmt.Println(string(jsonData2))

}
