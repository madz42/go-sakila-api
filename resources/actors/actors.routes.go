package actors

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {

	router := chi.NewRouter()
	router.Get("/", ListActors)
	router.Get("/{actorId}", GetActorById)
	router.Get("/search/{name}", GetActorByName)
	router.Post("/", CreateActor)
	router.Patch("/{actorId}", EditActorById)
	router.Delete("/{actorId}", DeleteActorById)

	return router
}
