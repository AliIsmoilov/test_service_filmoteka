package repository

import (
	"context"
	"errors"
	"fmt"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/constatnts"
	"time"

	"github.com/google/uuid"
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

// Update film with transaction
func (r *filmsRepo) Update(ctx context.Context, film *models.Film) (*models.Film, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return &models.Film{}, fmt.Errorf("error starting transaction: %w", tx.Error)
	}

	// Delete existing film actors
	if err := tx.Table("film_actors").
		Where("film_id = ?", film.ID).
		Delete(models.FilmActor{}).
		Error; err != nil {
		tx.Rollback()
		return &models.Film{}, fmt.Errorf("error deleting film actors: %w", err)
	}

	// Create new film actors
	if err := tx.Table("film_actors").
		Create(film.FilmActors).
		Error; err != nil {
		tx.Rollback()
		return &models.Film{}, fmt.Errorf("error creating film actors: %w", err)
	}

	// Update film record
	res := tx.Model(&models.Film{}).
		Where("id = ?", film.ID).
		Updates(map[string]interface{}{
			"title":        film.Title,
			"description":  film.Description,
			"release_date": film.ReleaseDate,
			"rating":       film.Rating,
			"updated_at":   time.Now(),
		})

	if res.Error != nil {
		tx.Rollback()
		return &models.Film{}, fmt.Errorf("error updating film: %w", res.Error)
	} else if res.RowsAffected == 0 {
		tx.Rollback()
		return &models.Film{}, fmt.Errorf("no rows affected in film update")
	}

	if err := tx.Commit().Error; err != nil {
		return &models.Film{}, fmt.Errorf("error committing transaction: %w", err)
	}

	return film, nil
}

// Delete film
func (r *filmsRepo) Delete(ctx context.Context, id uuid.UUID) error {

	res := r.db.Table("films").
		Where("id=?", id).
		Update("deleted_at", time.Now())
	if res.Error != nil {
		return fmt.Errorf("error in storage.Delete: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return fmt.Errorf("error in storage.Delete: %w", constatnts.ErrRowsAffectedZero)
	}

	return nil
}

// GetByID film
func (r *filmsRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Film, error) {

	resp := models.Film{}
	res := r.db.
		Table("films").
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Preload("FilmActors").
		Preload("FilmActors.Actor").
		First(&resp)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &models.Film{}, fmt.Errorf("film not found %w", constatnts.ErrRecordNotFound)
		}
		return &models.Film{}, res.Error
	}
	return &resp, nil
}
