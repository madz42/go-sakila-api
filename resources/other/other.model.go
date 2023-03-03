package other

import (
	"net/http"
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
	return nil
}

type FilmActorResponse struct {
	*FilmActor
}

func (f *FilmActorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
