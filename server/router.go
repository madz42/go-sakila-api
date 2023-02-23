package server

import (
	"go-sakila-api/resources/actors"
	"go-sakila-api/resources/films"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	router := chi.NewRouter()
	router.Mount("/actors", actors.Routes())
	router.Mount("/films", films.Routes())
	return router
}
