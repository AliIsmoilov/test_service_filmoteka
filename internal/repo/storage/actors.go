package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"test_service_filmoteka/internal/models"
	todos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/constatnts"
)

// Actor Repository
type actorsRepo struct {
	db *gorm.DB
}

// Actors Repository constructor
func NewActorsRepository(db *gorm.DB) todos.ActorsRepository {
	return &actorsRepo{db: db}
}

// Create actor
func (r *actorsRepo) Create(ctx context.Context, actor *models.Actor) (*models.Actor, error) {

	res := r.db.Model(models.Actor{}).Create(actor)
	if res.Error != nil {
		return &models.Actor{}, fmt.Errorf("error in storage.Create: %w", res.Error)
	}
	return actor, nil
}

// Update actor
func (r *actorsRepo) Update(ctx context.Context, actor *models.Actor) (*models.Actor, error) {

	res := r.db.Model(models.Actor{}).
		Where("id=?", actor.ID).
		Updates(actor)
	if res.Error != nil {
		return &models.Actor{}, fmt.Errorf("error in storage.Update: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return &models.Actor{}, fmt.Errorf("error in storage.Update: %w", constatnts.ErrRowsAffectedZero)
	}
	return &models.Actor{}, nil
}

// Delete actor
func (r *actorsRepo) Delete(ctx context.Context, id uuid.UUID) error {

	res := r.db.Table("actors").
		Where("id=?", id).
		Update("deleted_at", time.Now())
	if res.Error != nil {
		return fmt.Errorf("error in storage.Delete: %w", res.Error)
	} else if res.RowsAffected == 0 {
		return fmt.Errorf("error in storage.Delete: %w", constatnts.ErrRowsAffectedZero)
	}

	return nil
}

// GetByID actor
func (r *actorsRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.Actor, error) {

	resp := models.Actor{}
	res := r.db.
		Table("actors").
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		First(&resp)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &models.Actor{}, fmt.Errorf("actor not found %w", constatnts.ErrRecordNotFound)
		}
		return &models.Actor{}, res.Error
	}
	return &resp, nil
}

// GetAll actor
func (r *actorsRepo) GetAll(ctx context.Context, req models.ActorsListReq) (*models.ActorsListResp, error) {

	var actors []models.Actor
	var count int64

	tx := r.db.Table("actors").
		Where("deleted_at IS NULL")

	if req.Search != "" {
		req.Search = "%" + req.Search + "%"
		tx = tx.Where("name ilike ?", req.Search)
	}

	res := tx.Count(&count)
	if res.Error != nil {
		return &models.ActorsListResp{}, res.Error
	}
	if req.Page > 0 && req.Limit > 0 {
		offset := (req.Page - 1) * req.Limit
		tx = tx.Offset(int(offset)).
			Limit(int(req.Limit))
	} else if req.Limit > 0 {
		tx = tx.Limit(int(req.Limit))
	}

	tx = tx.Find(&actors)
	if tx.Error != nil {
		return &models.ActorsListResp{}, tx.Error
	}

	return &models.ActorsListResp{
		Count:  int(count),
		Actors: actors,
	}, nil
}
