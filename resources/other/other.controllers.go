package other

import (
	"encoding/json"
	"errors"
	db "go-sakila-api/database"
	e "go-sakila-api/error"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CreateRelFilmActor(w http.ResponseWriter, r *http.Request) {
	var newRel FilmActor
	err := json.NewDecoder(r.Body).Decode(&newRel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.DB.Create(&newRel)
	if result.Error != nil {
		// handle the error
	}
}

func DeleteRelFilmActor(w http.ResponseWriter, r *http.Request) {
	ActorId := chi.URLParam(r, "actorId")
	FilmId := chi.URLParam(r, "filmId")
	var oldRel FilmActor
	result := db.DB.Where("actor_id = ? AND film_id = ?", ActorId, FilmId).Find(&oldRel)
	if result.Error != nil {
		log.Println("error1")
		render.Render(w, r, e.ErrServerInternal(result.Error))
		return
	}
	if result.RowsAffected == 0 {
		log.Println("error2")
		render.Render(w, r, e.ErrNotFound(errors.New("no match")))
		return
	}
	db.DB.Where("actor_id = ? AND film_id = ?", oldRel.ActorId, oldRel.FilmId).Delete(&oldRel)
}
