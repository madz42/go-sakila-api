package server

import (
	"go-sakila-api/resources/actors"
	"go-sakila-api/resources/films"
	"go-sakila-api/resources/other"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func Router() chi.Router {
	router := chi.NewRouter()

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/actors", actors.Routes())
	router.Mount("/films", films.Routes())
	router.Mount("/other", other.Routes())
	return router
}
