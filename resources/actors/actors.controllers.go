package actors

import (
	"errors"
	db "go-sakila-api/database"
	e "go-sakila-api/error"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	db.DB.Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor
	if actor.FirstName != "" && actor.LastName != "" {
		var chkActor Actor
		result := db.DB.Where("first_name = ? AND last_name = ?", actor.FirstName, actor.LastName).First(&chkActor)
		if result.Error != nil {
			db.DB.Create(actor)
			render.Status(r, http.StatusCreated)
			render.Render(w, r, NewActorResponse(actor))
		} else {
			render.Render(w, r, e.ErrForbid(errors.New("actor already exists")))
		}
	} else {
		render.Render(w, r, e.ErrInvalidRequest(errors.New("missing name")))
	}
}

func DeleteActorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "actorId")
	var actor Actor
	result := db.DB.First(&actor, id)
	if result.Error != nil {
		log.Println("Delete actor by id: ", id, " - NOT FOUND")
		render.Render(w, r, e.ErrNotFound(result.Error))
		return
	}
	// db.DB.Exec("ALTER TABLE film_actor ADD CONSTRAINT fk_film_actor_actor FOREIGN KEY (actor_id) REFERENCES actor(actor_id) ON DELETE CASCADE")
	db.DB.Delete(&actor)
	log.Println("Delete actor by id: ", id, " - DELETED")
	render.Status(r, http.StatusNoContent)
}

func EditActorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "actorId")
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor
	var chkActor Actor
	result := db.DB.First(&chkActor, id)
	if result.Error != nil {
		log.Println("Edit actor by id: ", id, " - NOT FOUND")
		render.Render(w, r, e.ErrNotFound(result.Error))
		return
	}
	if actor.FirstName != "" {
		chkActor.FirstName = strings.ToUpper(actor.FirstName)
	}
	if actor.LastName != "" {
		chkActor.LastName = strings.ToUpper(actor.LastName)
	}
	chkActor.LastUpdate = time.Now()
	db.DB.Save(&chkActor)
	log.Println("Edit actor by id: ", id, " - EDITED")
	// render.Status(r, http.StatusNoContent)
	render.Render(w, r, NewActorResponse(&chkActor))
}

func GetActorById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "actorId")
	var actor Actor
	result := db.DB.First(&actor, id)
	if result.Error != nil {
		log.Println("Get actor by id: ", id, " - NOT FOUND")
		render.Render(w, r, e.ErrNotFound(result.Error))
		return
	}
	log.Println("Get actor by id: ", id, " - FOUND")
	render.Render(w, r, NewActorResponse(&actor))
}

func GetActorByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	var actors []*Actor
	result := db.DB.Where("first_name LIKE ? OR last_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&actors)
	if result.Error != nil {
		render.Render(w, r, e.ErrServerInternal(result.Error))
		return
	}
	log.Println("Search actor by '", name, "' - results:", result.RowsAffected)
	if result.RowsAffected == 0 {
		render.Render(w, r, e.ErrNotFound(errors.New("no match")))
		return
	}
	render.RenderList(w, r, NewActorListResponse(actors))
}
