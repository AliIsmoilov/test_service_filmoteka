package mappers

import (
	"test_service_filmoteka/internal/models"

	"github.com/google/uuid"
)

func ToFilm(req *models.FilmSwagger, id uuid.UUID) *models.Film {
	return &models.Film{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		ReleaseDate: req.ReleaseDate,
		Rating:      req.Rating,
		FilmActors:  ToFilmActors(req.ActorIds, id.String()),
	}
}

func ToFilmActors(actorIds []string, fimId string) []models.FilmActor {
	resp := []models.FilmActor{}
	for _, actorId := range actorIds {
		resp = append(resp, models.FilmActor{
			ID:      uuid.New(),
			FilmId:  fimId,
			ActorId: actorId,
		})
	}
	return resp
}
