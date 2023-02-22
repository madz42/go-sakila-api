package server

import (
	"go-sakila-api/resources/actors"

	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {
	router := chi.NewRouter()
	router.Mount("/actors", actors.Routes())
	return router
}
