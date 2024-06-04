package main

import (
	"errors"
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

// Divide handles the divide page requests
// It divides 100 by the provided value and returns the result
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

// divideValues divides two float32 numbers and returns the result
// If the divisor is 0, it returns an error
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("you cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the entry point for the application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// Functions that begin with a capital letter are accessible outside their package
// (not applicable for the main package)
