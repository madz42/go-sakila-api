package other

import (
	"net/http"
	// "github.com/go-chi/render"
)

type FilmActor struct {
	ActorId int `gorm:"type:smallint"`
	FilmId  int `gorm:"type:smallint"`
}

func (FilmActor) TableName() string {
	return "film_actor"
}

type FilmActorRequest struct {
	*FilmActor
}

func (a *FilmActorRequest) Bind(r *http.Request) error {
	// if a.ActorId == nil {
	// 	return errors.New("missing required Actor ID")
	// }
	// if a.FilmId == nil {
	// 	return errors.New("missing required Actor ID")
	// }

	// a.FilmActor.ActorId = strings.ToUpper(a.Actor.FirstName)
	// a.FilmActor.FilmId = strings.ToUpper(a.Actor.LastName)

	return nil
}

type FilmActorResponse struct {
	*FilmActor
}

// func NewFilmActorResponse(film *FilmActor) *FilmActorResponse {
// 	return &FilmActorResponse{
// 		Film:   film,
// 		Actors: actorResponses,
// 	}
// }

// func NewFilmListResponse(films []*Film) []render.Renderer {
// 	list := []render.Renderer{}
// 	for _, film := range films {
// 		list = append(list, NewFilmResponse(film))
// 	}
// 	return list
// }

func (f *FilmActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// type Actor struct {
// 	ActorId   int    `gorm:"type:smallint;primaryKey"`
// 	FirstName string `gorm:"type:varchar(45)"`
// 	LastName  string `gorm:"type:varchar(45)"`
// 	// LastUpdate time.Time `gorm:"autoCreateTime"`
// }

// func (Actor) TableName() string {
// 	return "actor"
// }

// type ActorResponse struct {
// 	*Actor
// }

// func NewActorResponse(actor *Actor) *ActorResponse {
// 	return &ActorResponse{actor}
// }

// func (a *ActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }
