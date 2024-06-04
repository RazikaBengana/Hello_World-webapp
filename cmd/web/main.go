package main

import (
	"fmt"
	"github.com/RazikaBengana/Hello_World-webapp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
