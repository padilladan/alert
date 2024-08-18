package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Contact(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Contact",
		"page":  "contact",
	}
	return c.Render(http.StatusOK, "contact.html", data)
}
