package server

import (
	"net/http"
	"strings"

	"test_service_filmoteka/docs"
	"test_service_filmoteka/pkg/csrf"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	repoHttp "test_service_filmoteka/internal/repo/delivery/http"
	repoRepository "test_service_filmoteka/internal/repo/storage"
	repoUseCase "test_service_filmoteka/internal/repo/usecase"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:8080
// @BasePath /api/v1
func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repositories
	cRepo := repoRepository.NewActorsRepository(s.db)
	commUC := repoUseCase.NewActorsUseCase(s.cfg, cRepo, s.logger)
	actorsHandlers := repoHttp.NewHandler(s.cfg, commUC, s.logger)

	filmRepo := repoRepository.NewFilmsRepository(s.db)
	filmUC := repoUseCase.NewFilmUseCase(s.cfg, filmRepo, s.logger)
	filmsHandlers := repoHttp.NewFilmHandler(s.cfg, filmUC, s.logger)

	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "App API"
	docs.SwaggerInfo.Description = "REST API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// e.Start(":5050")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID, csrf.CSRFHeader},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10, // 1 KB
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))
	e.Use(middleware.RequestID())

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	v1 := e.Group("/v1")

	health := v1.Group("/health")
	actorsGroup := v1.Group("/actors")
	repoHttp.MapActorsRoutes(actorsGroup, actorsHandlers)

	filmsGroup := v1.Group("/films")
	repoHttp.MapFilmRoutes(filmsGroup, filmsHandlers)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy!"})
	})

	return nil
}
