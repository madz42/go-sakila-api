package films

import (
	db "go-sakila-api/database"
	"net/http"

	"github.com/go-chi/render"
)

func ListFilms(w http.ResponseWriter, r *http.Request) {
	var films []*Film
	db.DB.Find(&films)
	render.RenderList(w, r, NewFilmListResponse(films))
}
