package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home handles the home page requests
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About handles the about page requests
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// addValues takes 2 integers and returns their sum
func addValues(x, y int) int {
	return x + y
}

// main is the entry point for the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starrting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// Functions that begin with a capital letter are accessible outside their package
// (not applicable for the main package)
