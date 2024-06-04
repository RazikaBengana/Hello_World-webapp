package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// Functions that begin with a capital letter are accessible outside their package
// (not applicable for the main package)
