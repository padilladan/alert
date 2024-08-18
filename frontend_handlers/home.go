package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Home(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Home",
		"page":  "home",
	}
	return c.Render(http.StatusOK, "home.html", data)
}
