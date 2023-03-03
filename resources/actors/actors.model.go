package actors

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/render"
)

type Actor struct {
	ActorId    int       `gorm:"type:smallint;primaryKey"`
	FirstName  string    `gorm:"type:varchar(45)"`
	LastName   string    `gorm:"type:varchar(45)"`
	LastUpdate time.Time `gorm:"autoCreateTime"`
	Films      []*Film   `gorm:"many2many:film_actor;joinForeignKey:actor_id;joinReferences:film_id"`
}

func (Actor) TableName() string {
	return "actor"
}

type ActorRequest struct {
	*Actor
}

func (a *ActorRequest) Bind(r *http.Request) error {
	if a.Actor == nil {
		return errors.New("missing required Actor fields")
	}

	a.Actor.FirstName = strings.ToUpper(a.Actor.FirstName)
	a.Actor.LastName = strings.ToUpper(a.Actor.LastName)

	return nil
}

type ActorResponse struct {
	*Actor
	Films []*FilmResponse `json:"Films"`
}

func NewActorResponse(actor *Actor) *ActorResponse {
	filmResponses := make([]*FilmResponse, 0, len(actor.Films))
	for _, film := range actor.Films {
		filmResponses = append(filmResponses, NewFilmResponse(film))
	}
	return &ActorResponse{
		Actor: actor,
		Films: filmResponses,
	}
}

func NewActorListResponse(actors []*Actor) []render.Renderer {
	list := []render.Renderer{}
	for _, actor := range actors {
		list = append(list, NewActorResponse(actor))
	}
	return list
}

func (a *ActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Film struct {
	FilmId      int    `gorm:"type:smallint;primaryKey"`
	Title       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	ReleaseYear int    `gorm:"type:year(4)"`

	Length int `gorm:"type:smallint"`

	Rating string `gorm:"type:enum('G','PG','PG-13','R','NC-17')"`
}

func (Film) TableName() string {
	return "film"
}

type FilmResponse struct {
	*Film
}

func NewFilmResponse(film *Film) *FilmResponse {
	return &FilmResponse{film}
}

func (f *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
