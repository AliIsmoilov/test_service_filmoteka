package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"test_service_filmoteka/config"
	"test_service_filmoteka/pkg/jwt"
	"test_service_filmoteka/pkg/logger"
	"test_service_filmoteka/pkg/utils"

	constants "test_service_filmoteka/pkg/constatnts"

	"github.com/labstack/echo/v4"
)

// Middleware manager
type MiddlewareManager struct {
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, origins: origins, logger: logger}
}

func (jwta *JWTRoleAuthorizer) Check() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return jwta.CheckAuth(next)
	}
}

// CORSWithConfig returns a CORS middleware with config.
// See: [CORS].
func (jwta *JWTRoleAuthorizer) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("CheckAuth........")
		req := c.Request()
		allowed, err := jwta.checkPermission(req)
		if err != nil {
			if errors.Is(err, constants.ErrAuthIncorrect) || errors.Is(err, constants.ErrAuthNotGiven) {
				utils.HandleForbiddenErrWithMessage(c.Response(), err, err.Error())
				// return err
			}

			if errors.Is(err, constants.InvalidToken) {
				utils.HandleUnauthorizedWithMessage(c.Response(), "Ошибка: для решение этой ошибки выйдите из аккаунта и зайдите обратно. Просим прошение за неудобства")
				return err
			}

			return echo.NewHTTPError(http.StatusBadGateway, err.Error())
		}
		if !allowed {
			utils.HandleUnauthorizedWithMessage(c.Response(), `Forbidden`)
			return err
		}

		return next(c)
	}
}

func (jwta *JWTRoleAuthorizer) checkPermission(r *http.Request) (bool, error) {
	path := r.URL.Path

	if strings.Contains(path, "swagger") {
		return true, nil
	}

	userMetadata, err := jwt.GetUserMetadata(constants.JWTSecretKey, r.Header.Get(constants.AuthorizationHeader))
	if err != nil {
		if errors.Is(err, constants.InvalidToken) {
			userMetadata, err = jwt.GetUserMetadata(constants.JWTKEY, r.Header.Get(constants.AuthorizationHeader))
			if err != nil {
				return false, err
			}
		} else {
			return false, err
		}
	}
	role := userMetadata["role"].(string)
	method := r.Method
	enforsed, err := jwta.enforcer.Enforce(role, path, method)
	if enforsed && role == "unauthorized" {
		// Authorize request with basicAuth
		username, password, ok := r.BasicAuth()
		if !ok {
			return false, constants.ErrAuthNotGiven
		}

		if !IsSuperAdmin(username, password) {
			return false, constants.ErrAuthIncorrect
		}
	}
	return enforsed, err
}
