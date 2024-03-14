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
type filmsUC struct {
	cfg       *config.Config
	filmsRepo repos.FilmsRepository
	logger    logger.Logger
}

// Actor UseCase constructor
func NewFilmUseCase(cfg *config.Config, filmsRepo repos.FilmsRepository, logger logger.Logger) repos.FilmsRepository {
	return &filmsUC{cfg: cfg, filmsRepo: filmsRepo, logger: logger}
}

// Create film
func (u *filmsUC) Create(ctx context.Context, film *models.Film) (*models.Film, error) {
	return u.filmsRepo.Create(ctx, film)
}

// GetAll films
func (u *filmsUC) GetAll(ctx context.Context, req models.FilmsListReq) (*models.FilmsListResp, error) {
	return u.filmsRepo.GetAll(ctx, req)
}

// Create film
func (u *filmsUC) Delete(ctx context.Context, id uuid.UUID) error {
	return u.filmsRepo.Delete(ctx, id)
}

// GetAll film
func (u *filmsUC) GetByID(ctx context.Context, id uuid.UUID) (*models.Film, error) {
	return u.filmsRepo.GetByID(ctx, id)
}

// Update film
func (u *filmsUC) Update(ctx context.Context, film *models.Film) (*models.Film, error) {
	return u.filmsRepo.Update(ctx, film)
}
