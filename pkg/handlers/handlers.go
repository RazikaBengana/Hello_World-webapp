package handlers

import (
	"github.com/RazikaBengana/Hello_World-webapp/pkg/render"
	"net/http"
)

// Home handles the home page requests
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About handles the about page requests
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
