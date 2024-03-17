package todos

import "github.com/labstack/echo/v4"

// HTTP Handlers interface
type ActorHandlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetActorFilms() echo.HandlerFunc
}

type FilmHandlers interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetFilmActors() echo.HandlerFunc
}

type UserHandlers interface {
	SignUp() echo.HandlerFunc
}
