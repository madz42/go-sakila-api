package films

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {

	router := chi.NewRouter()
	router.Get("/", ListFilms)
	router.Get("/{filmId}", GetFilmById)
	router.Get("/search/{name}", GetFilmByName)

	return router
}
