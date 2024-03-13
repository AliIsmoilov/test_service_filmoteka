package http

import (
	"github.com/labstack/echo/v4"

	todos "test_service_filmoteka/internal/repo"
)

// Map todos routes
func MapToDosRoutes(group *echo.Group, h todos.Handlers) {
	// docs.SwaggerInfo.Title = cfg.ServiceName
	// docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}
	group.POST("", h.Create())
	group.DELETE("/:id", h.Delete())
	group.PUT("/:id", h.Update())
	group.GET("/list", h.GetAll())
	group.GET("/:id", h.GetByID())
}
