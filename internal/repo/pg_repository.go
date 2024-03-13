//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package todos

import (
	"context"

	"test_service_filmoteka/internal/models"
	"test_service_filmoteka/pkg/utils"

	"github.com/google/uuid"
)

// Actor repository interface
type ActorsRepository interface {
	Create(ctx context.Context, blog *models.Actor) (*models.Actor, error)
	Update(ctx context.Context, todo *models.Actor) (*models.Actor, error)
	Delete(ctx context.Context, todoID uuid.UUID) error
	GetByID(ctx context.Context, blogID uuid.UUID) (*models.Actor, error)
	GetAll(ctx context.Context, title string, query *utils.PaginationQuery) (*models.ActorsList, error)

	// CreateNews(ctx context.Context, new *models.News) (*models.News, error)
}
