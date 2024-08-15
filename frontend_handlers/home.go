package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "home.tmpl", map[string]interface{}{
		"title": "Home Page",
	})
}
