package http

import (
	"net/http"
	"test_service_filmoteka/config"
	"test_service_filmoteka/internal/models"
	repos "test_service_filmoteka/internal/repo"
	"test_service_filmoteka/pkg/constatnts"
	"test_service_filmoteka/pkg/httpErrors"
	"test_service_filmoteka/pkg/jwt"
	"test_service_filmoteka/pkg/logger"
	"test_service_filmoteka/pkg/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Users handlers
type usersHandlers struct {
	cfg    *config.Config
	userUC repos.UsersUseCase
	logger logger.Logger
}

// NewBlogHandlers Actor handlers constructor
func NewUsersHandler(cfg *config.Config, userUC repos.UsersUseCase, logger logger.Logger) repos.UserHandlers {
	return &usersHandlers{cfg: cfg, userUC: userUC, logger: logger}
}

// SignUp
// @Summary user sign up
// @Description DESCRIPTION: 
// @Description for creating admin role and Get access to all apis signup with role 2
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body models.UserSwagger true "body"
// @Success 201 {object} models.SignUpSwagger
// @Failure 500 {object} httpErrors.RestErr
// @Router /users/ [post]
func (h *usersHandlers) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {

		reqBody := &models.User{}
		if err := utils.BodyParser(c.Request(), &reqBody); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
		}
		reqBody.ID = uuid.New()

		resp, err := h.userUC.SignUp(c.Request().Context(), reqBody)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		tokenCredentials := map[string]string{
			"id":           resp.ID.String(),
			"name":         resp.Name,
			"phone_number": resp.PhoneNumber,
			"role":         constatnts.Roles[resp.Role],
			"user_name":    resp.UserName,
			"password":     resp.Password,
		}

		// Generate a new pair of access and refresh tokens.
		tokens, err := jwt.GenerateNewUserTokens(resp.ID.String(), tokenCredentials)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, models.SignUpSwagger{
			Name:         resp.Name,
			PhoneNumber:  resp.PhoneNumber,
			Role:         resp.Role,
			UserName:     resp.UserName,
			Password:     resp.Password,
			AccessToken:  tokens.Access,
			RefreshToken: tokens.Refresh,
		})
	}
}
