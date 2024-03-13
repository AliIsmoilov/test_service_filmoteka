package models

import (
	"time"

	"github.com/google/uuid"
)

// Blog Swagger model
type ActorSwagger struct {
	Title string `json:"title" db:"title" validate:"required,gte=3"`
}

// All ToDo response
type ActorsList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Actors     []Actor `json:"actors"`
}

type Actor struct {
	ID        uuid.UUID `json:"id" gorm:"id" validate:"omitempty,uuid"`
	Name      string    `json:"name" gorm:"name" validate:"required,gte=3"`
	Gender    string    `json:"gender" gorm:"gender"`
	BirthDate string    `json:"birth_date" gorm:"birth_date"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"<-:update"`
	DeletedAt time.Time `gorm:"<-:update"`
}
