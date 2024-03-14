package repository

import (
	"context"
	"fmt"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"

	"gorm.io/gorm"
)

type filmsRepo struct {
	db *gorm.DB
}

// Actors Repository constructor
func NewFilmsRepository(db *gorm.DB) repos.FilmsRepository {
	return &filmsRepo{db: db}
}

// Create film
func (r *filmsRepo) Create(ctx context.Context, film *models.Film) (*models.Film, error) {
	res := r.db.Model(models.Film{}).Create(film)
	if res.Error != nil {
		return &models.Film{}, fmt.Errorf("error in storage.Create: %w", res.Error)
	}
	return film, nil
}

// GetAll film
func (r *filmsRepo) GetAll(ctx context.Context, req models.FilmsListReq) (*models.FilmsListResp, error) {

	var films []models.Film
	var count int64

	tx := r.db.Table("films").
		Where("deleted_at IS NULL")

	if req.Search != "" {
		req.Search = "%" + req.Search + "%"
		tx = tx.Where("title ilike ?", req.Search)
	}

	res := tx.Count(&count)
	if res.Error != nil {
		return &models.FilmsListResp{}, res.Error
	}
	if req.Page > 0 && req.Limit > 0 {
		offset := (req.Page - 1) * req.Limit
		tx = tx.Offset(int(offset)).
			Limit(int(req.Limit))
	} else if req.Limit > 0 {
		tx = tx.Limit(int(req.Limit))
	}

	tx = tx.Preload("FilmActors").Preload("FilmActors.Actor").Find(&films)
	if tx.Error != nil {
		return &models.FilmsListResp{}, tx.Error
	}

	return &models.FilmsListResp{
		Count: int(count),
		Films: films,
	}, nil
}
