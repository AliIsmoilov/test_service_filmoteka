package usecase

import (
	"context"

	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/logger"

	"github.com/google/uuid"
)

// Actor UseCase
type actorsUC struct {
	cfg        *config.Config
	actorsRepo repos.ActorsRepository
	logger     logger.Logger
}

// Actor UseCase constructor
func NewActorsUseCase(cfg *config.Config, actorsRepo repos.ActorsRepository, logger logger.Logger) repos.ActorsUseCase {
	return &actorsUC{cfg: cfg, actorsRepo: actorsRepo, logger: logger}
}

// Create actor
func (u *actorsUC) Create(ctx context.Context, actor *models.Actor) (*models.Actor, error) {
	return u.actorsRepo.Create(ctx, actor)
}

// Update actor
func (u *actorsUC) Update(ctx context.Context, actor *models.Actor) (*models.Actor, error) {
	updated, err := u.actorsRepo.Update(ctx, actor)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// Delete actor
func (u *actorsUC) Delete(ctx context.Context, id uuid.UUID) error {

	if err := u.actorsRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

// GetByID actor
func (u *actorsUC) GetByID(ctx context.Context, id uuid.UUID) (*models.Actor, error) {

	return u.actorsRepo.GetByID(ctx, id)
}

// GetAll actors
func (u *actorsUC) GetAll(ctx context.Context, req models.ActorsListReq) (*models.ActorsListResp, error) {
	return u.actorsRepo.GetAll(ctx, req)
}
