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

func LoginPost(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	// validate credentials
	if email == "user@example.com" && password == "password123" {
		return c.String(http.StatusOK, "Login successful!")
	}
	return c.String(http.StatusUnauthorized, "Invalid email or password.")
}
