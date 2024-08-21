package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Configuration(c echo.Context) error {
	data := map[string]interface{}{
		"page": "Configuration",
	}

	return c.Render(http.StatusOK, "configuration.html", data)
}
