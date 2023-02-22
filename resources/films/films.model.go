package models

import (
	"time"
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
}

func (Film) TableName() string {
	return "film"
}
