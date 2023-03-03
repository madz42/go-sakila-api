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

// import (
// 	"errors"
// 	db "go-sakila-api/database"
// 	e "go-sakila-api/error"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/render"
// 	"gorm.io/gorm"
// )

// func ListActors(w http.ResponseWriter, r *http.Request) {
// 	var actors []*Actor
// 	db.DB.Find(&actors)
// 	render.RenderList(w, r, NewActorListResponse(actors))
// }

func CreateRelFilmActor(w http.ResponseWriter, r *http.Request) {
	var newRel FilmActor
	err := json.NewDecoder(r.Body).Decode(&newRel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println(newRel.ActorId)
	// fmt.Println(newRel.FilmId)
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
	// log.Println("Search actor by '", name, "' - results:", result.RowsAffected)
	if result.RowsAffected == 0 {
		log.Println("error2")
		render.Render(w, r, e.ErrNotFound(errors.New("no match")))
		return
	}
	db.DB.Where("actor_id = ? AND film_id = ?", oldRel.ActorId, oldRel.FilmId).Delete(&oldRel)
	// render.RenderList(w, r, NewActorListResponse(actors))
}

// func CreateActor(w http.ResponseWriter, r *http.Request) {
// 	var data ActorRequest
// 	if err := render.Bind(r, &data); err != nil {
// 		render.Render(w, r, e.ErrInvalidRequest(err))
// 	}
// 	actor := data.Actor
// 	if actor.FirstName != "" && actor.LastName != "" {
// 		var chkActor Actor
// 		result := db.DB.Where("first_name = ? AND last_name = ?", actor.FirstName, actor.LastName).First(&chkActor)
// 		if result.Error != nil {
// 			db.DB.Create(actor)
// 			render.Status(r, http.StatusCreated)
// 			render.Render(w, r, NewActorResponse(actor))
// 		} else {
// 			render.Render(w, r, e.ErrForbid(errors.New("actor already exists")))
// 		}
// 	} else {
// 		render.Render(w, r, e.ErrInvalidRequest(errors.New("missing name")))
// 	}
// }

// func DeleteActorById(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "actorId")
// 	var actor Actor
// 	result := db.DB.First(&actor, id)
// 	if result.Error != nil {
// 		log.Println("Delete actor by id: ", id, " - NOT FOUND")
// 		render.Render(w, r, e.ErrNotFound(result.Error))
// 		return
// 	}
// 	// db.DB.Exec("ALTER TABLE film_actor ADD CONSTRAINT fk_film_actor_actor FOREIGN KEY (actor_id) REFERENCES actor(actor_id) ON DELETE CASCADE")
// 	db.DB.Delete(&actor)
// 	log.Println("Delete actor by id: ", id, " - DELETED")
// 	render.Status(r, http.StatusNoContent)
// }

// func EditActorById(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "actorId")
// 	var data ActorRequest
// 	if err := render.Bind(r, &data); err != nil {
// 		render.Render(w, r, e.ErrInvalidRequest(err))
// 	}

// 	actor := data.Actor
// 	var chkActor Actor
// 	result := db.DB.First(&chkActor, id)
// 	if result.Error != nil {
// 		log.Println("Edit actor by id: ", id, " - NOT FOUND")
// 		render.Render(w, r, e.ErrNotFound(result.Error))
// 		return
// 	}
// 	if actor.FirstName != "" {
// 		chkActor.FirstName = strings.ToUpper(actor.FirstName)
// 	}
// 	if actor.LastName != "" {
// 		chkActor.LastName = strings.ToUpper(actor.LastName)
// 	}
// 	chkActor.LastUpdate = time.Now()
// 	db.DB.Save(&chkActor)
// 	log.Println("Edit actor by id: ", id, " - EDITED")
// 	// render.Status(r, http.StatusNoContent)
// 	render.Render(w, r, NewActorResponse(&chkActor))
// }

// func GetActorById(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "actorId")
// 	_, err := strconv.Atoi(id)
// 	if err != nil {
// 		render.Render(w, r, e.ErrInvalidRequest(err))
// 		return
// 	}

// 	// Get the actor with the specified ID, including their films
// 	actor := &Actor{}
// 	result := db.DB.Preload("Films").First(actor, id)
// 	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 		render.Render(w, r, e.ErrNotFound(errors.New("no films found")))
// 		return
// 	}
// 	if result.Error != nil {
// 		render.Render(w, r, e.ErrServerInternal(result.Error))
// 		return
// 	}
// 	render.Render(w, r, NewActorResponse(actor))
// }

// func GetActorByName(w http.ResponseWriter, r *http.Request) {
// 	name := chi.URLParam(r, "name")
// 	var actors []*Actor
// 	result := db.DB.Where("first_name LIKE ? OR last_name LIKE ?", "%"+name+"%", "%"+name+"%").Find(&actors)
// 	if result.Error != nil {
// 		render.Render(w, r, e.ErrServerInternal(result.Error))
// 		return
// 	}
// 	log.Println("Search actor by '", name, "' - results:", result.RowsAffected)
// 	if result.RowsAffected == 0 {
// 		render.Render(w, r, e.ErrNotFound(errors.New("no match")))
// 		return
// 	}
// 	render.RenderList(w, r, NewActorListResponse(actors))
// }
