package main

import (
	"alert/frontend_handlers"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

// Render templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Global method
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	renderer := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Routes
	e.GET("/", frontend_handlers.Home)
	e.GET("/about", frontend_handlers.About)
	e.GET("/contact", frontend_handlers.Contact)

	// 404 handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == http.StatusNotFound {
				err := frontend_handlers.NotFound(c)
				if err != nil {
					return
				}
				return
			}
		}
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}

	e.Logger.Fatal(e.Start(":8081"))
}
