package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NotFound(c echo.Context) error {
	data := map[string]interface{}{
		"title":   "Page Not Found",
		"page":    "404",
		"message": "Sorry, the page you are looking for does not exist.",
	}
	return c.Render(http.StatusNotFound, "404.html", data)
}
