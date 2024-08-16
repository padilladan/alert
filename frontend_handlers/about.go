package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func About(c echo.Context) error {
	data := map[string]interface{}{
		"title": "About",
		"page":  "about",
	}
	return c.Render(http.StatusOK, "about.html", data)
}
