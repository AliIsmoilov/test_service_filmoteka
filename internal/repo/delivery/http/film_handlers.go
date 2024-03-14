package http

import (
	"errors"
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
// @Router /films [post]
func (h *filmHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		reqBody := &models.FilmSwagger{}
		if err := utils.BodyParser(c.Request(), &reqBody); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
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
// @Param page query int false "page number" Format(page)
// @Param limit query int false "number of elements per page" Format(limit)
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

		filmsList, err := h.filmsUs.GetAll(c.Request().Context(), models.FilmsListReq{
			Limit:  uint32(pq.Limit),
			Page:   uint32(pq.Page),
			Search: pq.Search,
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
