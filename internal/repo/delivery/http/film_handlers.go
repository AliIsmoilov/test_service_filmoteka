package http

import (
	"errors"
	"fmt"
	"net/http"
	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/constatnts"
	"test_service_filmoteka/pkg/httpErrors"
	"test_service_filmoteka/pkg/logger"
	"test_service_filmoteka/pkg/mappers"
	"test_service_filmoteka/pkg/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type filmHandlers struct {
	cfg     *config.Config
	filmsUs repos.FilmsRepository
	logger  logger.Logger
}

// NewBlogHandlers Actor handlers constructor
func NewFilmHandler(cfg *config.Config, filmsUs repos.FilmUseCase, logger logger.Logger) repos.FilmHandlers {
	return &filmHandlers{cfg: cfg, filmsUs: filmsUs, logger: logger}
}

// CreateFilm
// @Summary create new film
// @Description create new film
// @Tags Film
// @Accept  json
// @Produce  json
// @Param body body models.FilmSwagger true "body"
// @Success 201 {object} models.Film
// @Failure 500 {object} httpErrors.RestErr
// @Security ApiKeyAuth
// @Router /films [post]
func (h *filmHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		reqBody := &models.FilmSwagger{}
		if err := utils.BodyParser(c.Request(), &reqBody); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}
		if err := reqBody.Validate(); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		createdFilm, err := h.filmsUs.Create(c.Request().Context(), mappers.ToFilm(reqBody, uuid.New()))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdFilm)
	}
}

// GetAll
// @Summary Get Films
// @Description Get all films
// @Tags Film
// @Accept  json
// @Produce  json
// @Param search query string false "search by title"
// @Param search_by_actor query string false "search by actor"
// @Param page query int false "page number" Format(page)
// @Param limit query int false "number of elements per page" Format(limit)
// @Param order_by query string false "order by title, rating, release_date"
// @Success 200 {object} models.FilmsListResp
// @Failure 500 {object} httpErrors.RestErr
// @Router /films/list [get]
func (h *filmHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		orderBy := c.QueryParam("order_by")
		if orderBy != "" && !utils.IsIncludedInSlice(orderBy, constatnts.FilmOrderBy) {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(http.StatusBadRequest, fmt.Errorf("%v is invalid order_by type", orderBy).Error())
		}

		filmsList, err := h.filmsUs.GetAll(c.Request().Context(), models.FilmsListReq{
			Limit:         uint32(pq.Limit),
			Page:          uint32(pq.Page),
			Search:        pq.Search,
			SearchByActor: c.QueryParam("search_by_actor"),
			OrderBy:       orderBy,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, filmsList)
	}
}

// Update
// @Summary Update film
// @Description update film actor
// @Tags Film
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param body body models.FilmSwagger true "body"
// @Success 200 {object} models.FilmSwagger
// @Failure 500 {object} httpErrors.RestErr
// @Security ApiKeyAuth
// @Router /films/{id} [put]
func (h *filmHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		reqBody := &models.FilmSwagger{}
		if err := utils.BodyParser(c.Request(), &reqBody); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		updatedfilm, err := h.filmsUs.Update(c.Request().Context(), mappers.ToFilm(reqBody, id))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedfilm)
	}
}

// Delete
// @Summary Delete film
// @Description delete film
// @Tags Film
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Security ApiKeyAuth
// @Router /films/{id} [delete]
func (h *filmHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.filmsUs.Delete(c.Request().Context(), id); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Get film
// @Description Get film by id
// @Tags Film
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Film
// @Failure 500 {object} httpErrors.RestErr
// @Router /films/{id} [get]
func (h *filmHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		actor, err := h.filmsUs.GetByID(c.Request().Context(), blogsID)
		if err != nil {
			if errors.Is(err, constatnts.ErrRecordNotFound) {
				utils.LogResponseError(c, h.logger, err)
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, actor)
	}
}

// GetFilmActors
// @Summary Get film actors
// @Description Get film actors by film_id
// @Tags Film
// @Accept  json
// @Produce  json
// @Param film_id path string true "film_id"
// @Success 200 {object} models.GetFilmActorsResp
// @Failure 500 {object} httpErrors.RestErr
// @Router /films/{film_id} [get]
func (h *filmHandlers) GetFilmActors() echo.HandlerFunc {
	return func(c echo.Context) error {

		filmID, err := uuid.Parse(c.Param("film_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}
		if !utils.IsUUID(filmID.String()) {
			utils.LogResponseError(c, h.logger, fmt.Errorf("%v is not valid uuid", filmID))
			return c.JSON(http.StatusBadRequest, fmt.Errorf("%v is not valid uuid", filmID).Error())
		}

		resp, err := h.filmsUs.GetFilmActors(c.Request().Context(), filmID)
		if err != nil {
			if errors.Is(err, constatnts.ErrRecordNotFound) {
				utils.LogResponseError(c, h.logger, err)
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, mappers.ToActorFromFilmActor(resp))
	}
}
