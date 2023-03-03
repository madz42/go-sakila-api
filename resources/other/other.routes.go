package other

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {

	router := chi.NewRouter()

	router.Post("/", CreateRelFilmActor)
	router.Delete("/{actorId}-{filmId}", DeleteRelFilmActor)

	return router
}
