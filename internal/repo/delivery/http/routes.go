package http

import (
	"github.com/labstack/echo/v4"

	repos "test_service_filmoteka/internal/repo"
)

// Map actors routes
func MapActorsRoutes(group *echo.Group, h repos.ActorHandlers) {
	// docs.SwaggerInfo.Title = cfg.ServiceName
	// docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}
	group.POST("", h.Create())
	group.DELETE("/:id", h.Delete())
	group.PUT("/:id", h.Update())
	group.GET("/list", h.GetAll())
	group.GET("/:id", h.GetByID())

}

// Map actors routes
func MapFilmRoutes(group *echo.Group, h repos.FilmHandlers) {
	group.POST("", h.Create())
	group.GET("/list", h.GetAll())
	group.DELETE("/:id", h.Delete())
	group.PUT("/:id", h.Update())
	group.GET("/:id", h.GetByID())
}
