package repository

import (
	"context"
	"database/sql"
	"test_service_filmoteka/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/test-go/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestCreateActor(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	actorRepo := NewActorsRepository(db)

	actor := &models.Actor{
		ID:     uuid.New(),
		Name:   gofakeit.Name(),
		Gender: "male",
	}
	actorRepo.Create(context.TODO(), actor)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestUpdateActor(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	actorRepo := NewActorsRepository(db)

	actor := &models.Actor{
		ID:     uuid.New(),
		Name:   gofakeit.Name(),
		Gender: "male",
	}
	actorRepo.Update(context.TODO(), actor)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteActor(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	actorRepo := NewActorsRepository(db)

	id := uuid.New()
	actorRepo.Delete(context.TODO(), id)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetByIDActor(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	actorRepo := NewActorsRepository(db)

	id := uuid.New()
	actorRepo.GetByID(context.TODO(), id)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetAllActor(t *testing.T) {
	sqlDB, db, mock := DbMock(t)
	defer sqlDB.Close()
	actorRepo := NewActorsRepository(db)

	req := models.ActorsListReq{}
	actorRepo.GetAll(context.TODO(), req)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func DbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqldb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}
	return sqldb, gormdb, mock
}
