package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a specified template file
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create a template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// Create a new buffer to hold the rendered template
	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache creates a cache for templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// Add the template set to the cache
		myCache[name] = ts
	}

	// Return the populated template cache
	return myCache, nil
}
