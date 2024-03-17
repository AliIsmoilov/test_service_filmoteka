package middleware

import (
	"fmt"
	"test_service_filmoteka/config"
	"test_service_filmoteka/pkg/logger"

	// constants"test_service_filmoteka/pkg/constatnts"

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

// func Check() error {
// 	fmt.Println("I have checked......")
// 	return nil
// }

func Check() echo.MiddlewareFunc {
	return CheckAuth()
}

// CORSWithConfig returns a CORS middleware with config.
// See: [CORS].
func CheckAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// req := c.Request()
			fmt.Println("CheckAuth........")
			return next(c)
		}
	}
}
