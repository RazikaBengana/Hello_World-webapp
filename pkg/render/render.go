package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplateTest renders a specified template file
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
	}
}

// tc = template cache
var tc = make(map[string]*template.Template)

// RenderTemplate renders a template, using cache if available
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check if template is in cache
	_, inMap := tc[t]
	if !inMap {
		// Create and cache the template if not found
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// Use cached template
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache parses and caches the template
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// Parse the template files
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add parsed template to cache
	tc[t] = tmpl

	return nil
}
