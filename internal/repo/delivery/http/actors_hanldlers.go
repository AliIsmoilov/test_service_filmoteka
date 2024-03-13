package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/httpErrors"
	"test_service_filmoteka/pkg/logger"
	"test_service_filmoteka/pkg/utils"
)

// Actor handlers
type actorHandlers struct {
	cfg      *config.Config
	actorsUC repos.ActorsUseCase
	logger   logger.Logger
}

// NewBlogHandlers Actor handlers constructor
func NewHandler(cfg *config.Config, actorsUC repos.ActorsUseCase, logger logger.Logger) repos.Handlers {
	return &actorHandlers{cfg: cfg, actorsUC: actorsUC, logger: logger}
}

// CreateActor
// @Summary create new actor
// @Description create new actor
// @Tags Actor
// @Accept  json
// @Produce  json
// @Param body body models.ActorSwagger true "body"
// @Success 201 {object} models.Actor
// @Failure 500 {object} httpErrors.RestErr
// @Router /actors [post]
func (h *actorHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		actor := &models.Actor{}
		if err := utils.SanitizeRequest(c, actor); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}

		actor.ID = uuid.New()
		createdActor, err := h.actorsUC.Create(c.Request().Context(), actor)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdActor)
	}
}

// Update
// @Summary Update actor
// @Description update new actor
// @Tags Actor
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param body body models.ActorSwagger true "body"
// @Success 200 {object} models.ActorSwagger
// @Failure 500 {object} httpErrors.RestErr
// @Router /actors/{id} [put]
func (h *actorHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		comm := &models.Actor{}
		if err = utils.SanitizeRequest(c, comm); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		updatedToDo, err := h.actorsUC.Update(c.Request().Context(), &models.Actor{
			ID:        id,
			Name:      comm.Name,
			Gender:    comm.Gender,
			BirthDate: comm.BirthDate,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, updatedToDo)
	}
}

// Delete
// @Summary Delete actor
// @Description delete actor
// @Tags Actor
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Router /actors/{id} [delete]
func (h *actorHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		if err = h.actorsUC.Delete(c.Request().Context(), blogsID); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Get actor
// @Description Get actor by id
// @Tags Actor
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Actor
// @Failure 500 {object} httpErrors.RestErr
// @Router /actors/{id} [get]
func (h *actorHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		blogsID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		actor, err := h.actorsUC.GetByID(c.Request().Context(), blogsID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, actor)
	}
}

// GetAll
// @Summary Get Actor
// @Description Get all actor
// @Tags Actor
// @Accept  json
// @Produce  json
// @Param search query string false "search by name"
// @Param page query int false "page number" Format(page)
// @Param limit query int false "number of elements per page" Format(limit)
// @Success 200 {object} models.ActorsListResp
// @Failure 500 {object} httpErrors.RestErr
// @Router /actors/list [get]
func (h *actorHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		toDoList, err := h.actorsUC.GetAll(c.Request().Context(), models.ActorsListReq{
			Limit:  uint32(pq.Limit),
			Page:   uint32(pq.Page),
			Search: pq.Search,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, toDoList)
	}
}
