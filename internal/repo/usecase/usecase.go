package usecase

import (
	"context"

	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/logger"
	"test_service_filmoteka/pkg/utils"

	"github.com/google/uuid"
)

// ToDos UseCase
type actorsUC struct {
	cfg        *config.Config
	actorsRepo repos.ActorsRepository
	logger     logger.Logger
}

// ToDos UseCase constructor
func NewToDosUseCase(cfg *config.Config, actorsRepo repos.ActorsRepository, logger logger.Logger) repos.ActorsUseCase {
	return &actorsUC{cfg: cfg, actorsRepo: actorsRepo, logger: logger}
}

// Create todo
func (u *actorsUC) Create(ctx context.Context, blog *models.Actor) (*models.Actor, error) {
	return u.actorsRepo.Create(ctx, blog)
}

// Update todo
func (u *actorsUC) Update(ctx context.Context, todo *models.Actor) (*models.Actor, error) {
	updatedToDo, err := u.actorsRepo.Update(ctx, todo)
	if err != nil {
		return nil, err
	}

	return updatedToDo, nil
}

// Delete todo
func (u *actorsUC) Delete(ctx context.Context, todoID uuid.UUID) error {

	if err := u.actorsRepo.Delete(ctx, todoID); err != nil {
		return err
	}

	return nil
}

// GetByID todo
func (u *actorsUC) GetByID(ctx context.Context, blogID uuid.UUID) (*models.Actor, error) {

	return u.actorsRepo.GetByID(ctx, blogID)
}

// GetAll todos
func (u *actorsUC) GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.ActorsList, error) {
	return u.actorsRepo.GetAll(ctx, title, query)
}
