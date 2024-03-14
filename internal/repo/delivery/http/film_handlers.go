package http

import (
	"net/http"
	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/httpErrors"
	"test_service_filmoteka/pkg/logger"
	"test_service_filmoteka/pkg/mappers"
	"test_service_filmoteka/pkg/utils"

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

		createdFilm, err := h.filmsUs.Create(c.Request().Context(), mappers.ToFilm(reqBody))
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
