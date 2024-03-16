package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// TODO: add janr
type Film struct {
	ID          uuid.UUID   `json:"id" gorm:"id" validate:"omitempty,uuid"`
	Title       string      `json:"title" gorm:"title"`
	Description string      `json:"description" gorm:"description"`
	ReleaseDate string      `json:"release_date" gorm:"release_date"`
	Rating      uint32      `json:"rating" gorm:"rating"`
	FilmActors  []FilmActor `json:"film_actors"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"<-:update"`
	DeletedAt   time.Time   `json:"deleted_at" gorm:"<-:update"`
}

type FilmActor struct {
	ID        uuid.UUID `json:"id" gorm:"id" validate:"omitempty,uuid"`
	FilmId    string    `json:"film_id"`
	ActorId   string    `json:"actor_id"`
	Actor     Actor     `json:"actor" gorm:"foreignKey:ActorId"`
	Film      Film      `json:"film" gorm:"foreignKey:FilmId"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"<-:update"`
}

type FilmSwagger struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ReleaseDate string   `json:"release_date"`
	Rating      uint32   `json:"rating"`
	ActorIds    []string `json:"actor_ids"`
}

func (f FilmSwagger) Validate() error {
	if len(f.Title) < 3 || len(f.Title) > 150 {
		return fmt.Errorf("length of film name should be between 3 and 150")
	}
	if len(f.Description) > 1000 {
		return fmt.Errorf("length of film description should not be more than 1000 sysmbols")
	}
	return nil
}

type FilmsListReq struct {
	Limit         uint32
	Page          uint32
	Search        string
	OrderBy       string
	SearchByActor string
}

type FilmsListResp struct {
	Count int    `json:"count"`
	Films []Film `json:"films"`
}

type GetFilmActorsResp struct {
	Count      int     `json:"count"`
	FilmActors []Actor `json:"film_actors"`
}

type GetActorFilmsResp struct {
	Count      int    `json:"count"`
	ActorFilms []Film `json:"actor_films"`
}
