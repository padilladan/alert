package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Alerts(c echo.Context) error {
	return c.Render(http.StatusOK, "alerts.tmpl", map[string]interface{}{
		"title": "Alert Page",
	})
}
