package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Alerts(c echo.Context) error {
	data := map[string]interface{}{
		"page": "Alerts",
	}

	return c.Render(http.StatusOK, "alerts.html", data)
}
