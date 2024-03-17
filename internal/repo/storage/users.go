package repository

import (
	"context"
	"fmt"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"

	"gorm.io/gorm"
)

type usersRepo struct {
	db *gorm.DB
}

// Actors Repository constructor
func NewUsersRepository(db *gorm.DB) repos.UsersRepository {
	return &usersRepo{db: db}
}

// Create actor
func (r *usersRepo) SignUp(ctx context.Context, user *models.User) (*models.User, error) {

	res := r.db.Model(models.User{}).Create(user)
	if res.Error != nil {
		return &models.User{}, fmt.Errorf("error in storage.Create: %w", res.Error)
	}
	return user, nil
}
