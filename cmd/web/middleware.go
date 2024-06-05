package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

/*
// WriteToConsole logs each request to the console before passing it to the next handler
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}
*/

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 // Ensure the cookie is not accessible via JavaScript
		Path:     "/",                  // Set the cookie to be accessible site-wide
		Secure:   false,                // Indicate whether the cookie should only be sent over HTTPS
		SameSite: http.SameSiteLaxMode, // Provide some protection against CSRF attacks
	})
	return csrfHandler
}
