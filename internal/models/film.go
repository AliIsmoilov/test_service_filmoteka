package models

import (
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

type FilmsListReq struct {
	Limit  uint32
	Page   uint32
	Search string
}

type FilmsListResp struct {
	Count int    `json:"count"`
	Films []Film `json:"films"`
}
