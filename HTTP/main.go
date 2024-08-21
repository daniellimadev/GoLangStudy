package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// URL to which you want to send the request
	url := "https://google.com"

	// Make the GET request
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	// Read the response body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// Print the response body
	fmt.Println(string(body))
}
