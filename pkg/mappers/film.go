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

func ToActorFromFilmActor(req []models.FilmActor) models.GetFilmActorsResp {
	resp := models.GetFilmActorsResp{Count: len(req)}
	for _, actor := range req {
		resp.FilmActors = append(resp.FilmActors, models.Actor{
			ID:        actor.ID,
			Name:      actor.Actor.Name,
			Gender:    actor.Actor.Gender,
			BirthDate: actor.Actor.BirthDate,
		})
	}
	return resp
}

func ToFilmFromFilmActor(req []models.FilmActor) models.GetActorFilmsResp {
	resp := models.GetActorFilmsResp{Count: len(req)}
	for _, film := range req {
		resp.ActorFilms = append(resp.ActorFilms, models.Film{
			ID:          film.ID,
			Title:       film.Film.Title,
			Description: film.Film.Description,
			ReleaseDate: film.Film.ReleaseDate,
			Rating:      film.Film.Rating,
		})
	}
	return resp
}
