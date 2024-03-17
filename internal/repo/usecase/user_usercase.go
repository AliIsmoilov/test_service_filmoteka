package usecase

import (
	"context"
	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/logger"
)

type usersUC struct {
	cfg       *config.Config
	usersRepo repos.UsersRepository
	logger    logger.Logger
}

// Actor UseCase constructor
func NewUsersUseCase(cfg *config.Config, usersRepo repos.UsersRepository, logger logger.Logger) repos.UsersUseCase {
	return &usersUC{cfg: cfg, usersRepo: usersRepo, logger: logger}
}

// Create actor
func (u *usersUC) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
	return u.usersRepo.SignUp(ctx, user)
}
