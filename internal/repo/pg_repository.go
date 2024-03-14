//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package todos

import (
	"context"

	"test_service_filmoteka/internal/models"

	"github.com/google/uuid"
)

// Actor repository interface
type ActorsRepository interface {
	Create(ctx context.Context, actor *models.Actor) (*models.Actor, error)
	Update(ctx context.Context, actor *models.Actor) (*models.Actor, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Actor, error)
	GetAll(ctx context.Context, req models.ActorsListReq) (*models.ActorsListResp, error)

	// CreateNews(ctx context.Context, new *models.News) (*models.News, error)
}

type FilmsRepository interface {
	Create(ctx context.Context, film *models.Film) (*models.Film, error)
	GetAll(ctx context.Context, req models.FilmsListReq) (*models.FilmsListResp, error)
	// Update(ctx context.Context, actor *models.Actor) (*models.Actor, error)
	// Delete(ctx context.Context, id uuid.UUID) error
	// GetByID(ctx context.Context, id uuid.UUID) (*models.Actor, error)

	// CreateNews(ctx context.Context, new *models.News) (*models.News, error)
}
