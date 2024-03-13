package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"test_service_filmoteka/internal/models"
	todos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/utils"
)

// Actor Repository
type actorsRepo struct {
	db *gorm.DB
}

// Actors Repository constructor
func NewActorsRepository(db *gorm.DB) todos.ActorsRepository {
	return &actorsRepo{db: db}
}

// Create todo
func (r *actorsRepo) Create(ctx context.Context, todo *models.Actor) (*models.Actor, error) {

	return &models.Actor{}, nil
}

// Update blog
func (r *actorsRepo) Update(ctx context.Context, blog *models.Actor) (*models.Actor, error) {

	return &models.Actor{}, nil
}

// Delete blog
func (r *actorsRepo) Delete(ctx context.Context, blogID uuid.UUID) error {

	return nil
}

// GetByID blog
func (r *actorsRepo) GetByID(ctx context.Context, blogId uuid.UUID) (*models.Actor, error) {

	return &models.Actor{}, nil
}

// GetAll ToDos
func (r *actorsRepo) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.ActorsList, error) {

	return &models.ActorsList{}, nil
}
