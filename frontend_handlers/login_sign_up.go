package frontend_handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginSignUp(c echo.Context) error {
	data := map[string]interface{}{
		"title": "Login & Sign Up Page",
		"page":  "login_sign_up",
	}
	return c.Render(http.StatusOK, "login_sign_up.html", data)
}
