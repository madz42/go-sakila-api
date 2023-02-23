package films

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type Film struct {
	FilmId          int       `gorm:"type:smallint;primaryKey"`
	Title           string    `gorm:"type:varchar(255)"`
	Description     string    `gorm:"type:text"`
	ReleaseYear     int       `gorm:"type:year"`
	LanguageId      int       `gorm:"type:tinyint"`
	RentalDuration  int       `gorm:"type:smallint"`
	RentalRate      float32   `gorm:"type:decimal(4,2)"`
	Length          int       `gorm:"type:smallint"`
	ReplacementCost float32   `gorm:"type:decimal(5,2)"`
	Rating          string    `gorm:"type:enum('G','PG','PG-13','R','NC-17')"`
	SpecialFeatures string    `gorm:"type:set('Trailers','Commentaries','Deleted Scenes','Behind the Scenes')"`
	LastUpdate      time.Time `gorm:"autoCreateTime"`
	Actors          []*Actor  `gorm:"many2many:film_actor;joinForeignKey:film_id;joinReferences:actor_id"`
}

func (Film) TableName() string {
	return "film"
}

type FilmRequest struct {
	*Film
}

//add Film Bind

type FilmResponse struct {
	*Film
	Actors []*ActorResponse `json:"Actors"`
}

func NewFilmResponse(film *Film) *FilmResponse {
	actorResponses := make([]*ActorResponse, 0, len(film.Actors))
	for _, actor := range film.Actors {
		actorResponses = append(actorResponses, NewActorResponse(actor))
	}
	return &FilmResponse{
		Film:   film,
		Actors: actorResponses,
	}
}

func NewFilmListResponse(films []*Film) []render.Renderer {
	list := []render.Renderer{}
	for _, film := range films {
		list = append(list, NewFilmResponse(film))
	}
	return list
}

func (f *FilmResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Actor struct {
	ActorId   int    `gorm:"type:smallint;primaryKey"`
	FirstName string `gorm:"type:varchar(45)"`
	LastName  string `gorm:"type:varchar(45)"`
	// LastUpdate time.Time `gorm:"autoCreateTime"`
}

func (Actor) TableName() string {
	return "actor"
}

type ActorResponse struct {
	*Actor
}

func NewActorResponse(actor *Actor) *ActorResponse {
	return &ActorResponse{actor}
}

func (a *ActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
