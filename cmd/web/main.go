package main

import (
	"fmt"
	"github.com/RazikaBengana/Hello_World-webapp/pkg/config"
	"github.com/RazikaBengana/Hello_World-webapp/pkg/handlers"
	"github.com/RazikaBengana/Hello_World-webapp/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point for the application
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
