package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// set up some middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// defiine application routes
	mux.Get("/", app.HomePage)

	return mux
}
