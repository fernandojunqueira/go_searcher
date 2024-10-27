package echo

import (
	"go_searcher/participant"
	programapocalipse "go_searcher/program_apocalipse"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handlers(pService participant.UseCase, cService programapocalipse.UseCase) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.GET("/participants", GetByName(pService))
	e.GET("/checkin/:id", Checkin(cService, pService))
	e.GET("/list", List(pService))
	e.PATCH("/clean", Clean(pService))
	return e
}

func Clean(s participant.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		s.Clean()
		return c.JSON(http.StatusOK, "pp")
	}
}

func List(s participant.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		pp, err := s.List()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, pp)
	}
}

func Checkin(s programapocalipse.UseCase, p participant.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		_, err := s.Get(id)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		pp, err := p.Update(id)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, pp)
	}
}

func GetByName(s participant.UseCase) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")

		if name == "" {
			return c.String(http.StatusBadRequest, "Missing 'name' query parameter")
		}

		p, err := s.Search(name)
		if err != nil && err.Error() == "erro buscando person do repositório: not found" {
			return c.String(http.StatusNotFound, err.Error())
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, p)
	}
}
