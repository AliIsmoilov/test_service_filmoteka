package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"id" validate:"omitempty,uuid"`
	Name        string    `json:"name" gorm:"name" validate:"required,gte=3"`
	PhoneNumber string    `json:"phone_number" gorm:"phone_number"`
	Role        int       `json:"role" gorm:"role"`
	UserName    string    `json:"user_name" gorm:"user_name"`
	Password    string    `json:"password" gorm:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"<-:update"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"<-:update"`
}

type UserSwagger struct {
	Name        string `json:"name" gorm:"name" validate:"required,gte=3"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
	Role        int    `json:"role" gorm:"role"`
	UserName    string `json:"user_name" gorm:"user_name"`
	Password    string `json:"password" gorm:"password"`
}

type SignUpSwagger struct {
	Name         string `json:"name" gorm:"name" validate:"required,gte=3"`
	PhoneNumber  string `json:"phone_number" gorm:"phone_number"`
	Role         int    `json:"role" gorm:"role"`
	UserName     string `json:"user_name" gorm:"user_name"`
	Password     string `json:"password" gorm:"password"`
	AccessToken  string `json:"acces_token"`
	RefreshToken string `json:"refresh_token"`
}
