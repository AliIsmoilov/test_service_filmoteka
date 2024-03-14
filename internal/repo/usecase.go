//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package todos

import (
	"context"

	"test_service_filmoteka/internal/models"

	"github.com/google/uuid"
)

// actors use case
type ActorsUseCase interface {
	Create(ctx context.Context, blog *models.Actor) (*models.Actor, error)
	Update(ctx context.Context, blog *models.Actor) (*models.Actor, error)
	Delete(ctx context.Context, blogID uuid.UUID) error
	GetByID(ctx context.Context, blogID uuid.UUID) (*models.Actor, error)
	GetAll(ctx context.Context, req models.ActorsListReq) (*models.ActorsListResp, error)
}

// actors use case
type FilmUseCase interface {
	Create(ctx context.Context, film *models.Film) (*models.Film, error)
	GetAll(ctx context.Context, req models.FilmsListReq) (*models.FilmsListResp, error)
	// Update(ctx context.Context, blog *models.Actor) (*models.Actor, error)
	// Delete(ctx context.Context, blogID uuid.UUID) error
	// GetByID(ctx context.Context, blogID uuid.UUID) (*models.Actor, error)
}
