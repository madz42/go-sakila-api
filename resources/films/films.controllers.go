package films

import (
	"errors"
	db "go-sakila-api/database"
	e "go-sakila-api/error"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	db.DB.Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "filmId")

	film := &Film{}
	result := db.DB.Preload("Actors").First(film, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		render.Render(w, r, e.ErrNotFound(errors.New("no films found")))
		return
	}
	if result.Error != nil {
		render.Render(w, r, e.ErrServerInternal(result.Error))
		return
	}

	// Render the actor with films as the response
	render.Render(w, r, NewFilmResponse(film))

	// var actor Actor
	// result := db.DB.First(&actor, id)
	// if result.Error != nil {
	// 	log.Println("Get actor by id: ", id, " - NOT FOUND")
	// 	render.Render(w, r, e.ErrNotFound(result.Error))
	// 	return
	// }
	// var films []*Film

	// log.Println("Get actor by id: ", id, " - FOUND")
	// render.Render(w, r, NewActorResponse(&actor))
}
