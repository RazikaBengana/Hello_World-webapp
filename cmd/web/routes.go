package main

import (
	"github.com/RazikaBengana/Hello_World-webapp/pkg/config"
	"github.com/RazikaBengana/Hello_World-webapp/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// routes sets up the application routes using the chi router and handlers
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
