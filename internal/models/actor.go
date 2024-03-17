package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Actor Swagger model
type ActorSwagger struct {
	Name      string `json:"name" gorm:"name" validate:"required,gte=3"`
	Gender    string `json:"gender" gorm:"gender"`
	BirthDate string `json:"birth_date" gorm:"birth_date"`
}

// All ToDo response
type ActorsListResp struct {
	Count  int     `json:"count"`
	Actors []Actor `json:"actors"`
}

type Actor struct {
	ID        uuid.UUID `json:"id" gorm:"id" validate:"omitempty,uuid"`
	Name      string    `json:"name" gorm:"name" validate:"required,gte=3"`
	Gender    string    `json:"gender" gorm:"gender"`
	BirthDate string    `json:"birth_date" gorm:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"<-:update"`
	DeletedAt time.Time `json:"deleted_at" gorm:"<-:update"`
}

type ActorsListReq struct {
	Limit  uint32
	Page   uint32
	Search string
}

func (a Actor) Validate() error {

	if a.Gender != "male" && a.Gender != "female" {
		return fmt.Errorf("gender should be male or female")
	}
	return nil
}
