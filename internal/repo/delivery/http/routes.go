package http

import (
	"github.com/labstack/echo/v4"

	"test_service_filmoteka/internal/repo"
)

// Map todos routes
func MapToDosRoutes(todoGroup *echo.Group, h todos.Handlers) {
	// docs.SwaggerInfo.Title = cfg.ServiceName
	// docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}
	todoGroup.POST("", h.Create())
	todoGroup.DELETE("/:id", h.Delete())
	todoGroup.PUT("/:id", h.Update())
	todoGroup.GET("/list", h.GetAll())
	todoGroup.GET("/:id", h.GetByID())
}
