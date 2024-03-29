package films

import (
	"errors"
	db "go-sakila-api/database"
	e "go-sakila-api/error"
	"log"
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

	render.Render(w, r, NewFilmResponse(film))
}

func GetFilmByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	var films []*Film
	result := db.DB.Where("title LIKE ?", "%"+name+"%").Find(&films)
	if result.Error != nil {
		render.Render(w, r, e.ErrServerInternal(result.Error))
		return
	}
	log.Println("Search film by '", name, "' - results:", result.RowsAffected)
	if result.RowsAffected == 0 {
		render.Render(w, r, e.ErrNotFound(errors.New("no match")))
		return
	}
	render.RenderList(w, r, NewFilmListResponse(films))
}
