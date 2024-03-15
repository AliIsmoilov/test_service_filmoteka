package repository

import (
	"context"
	"test_service_filmoteka/internal/models"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/test-go/testify/assert"
)

func TestCreateFilm(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	filmRepo := NewFilmsRepository(db)

	film := &models.Film{
		ID:          uuid.New(),
		Title:       gofakeit.Name(),
		Description: gofakeit.Name(),
	}
	filmRepo.Create(context.TODO(), film)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateFilm(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	filmRepo := NewFilmsRepository(db)
	film := &models.Film{
		ID:          uuid.New(),
		Title:       gofakeit.Name(),
		Description: gofakeit.Name(),
	}
	filmRepo.Update(context.TODO(), film)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteFilm(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	filmRepo := NewFilmsRepository(db)

	id := uuid.New()
	filmRepo.Delete(context.TODO(), id)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetByIDFilm(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	filmRepo := NewFilmsRepository(db)

	id := uuid.New()
	filmRepo.GetByID(context.TODO(), id)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetAllFilm(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	filmRepo := NewFilmsRepository(db)

	req := models.FilmsListReq{}
	filmRepo.GetAll(context.TODO(), req)
	assert.Nil(t, mock.ExpectationsWereMet())
}
