package main

import (
	"alert/frontend_handlers"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

// render template
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	//map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	//Initialize
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.tmpl")),
	}
	e.Renderer = renderer

	// routes
	e.GET("/", frontend_handlers.Home)
	e.GET("/another", frontend_handlers.Alerts)

	// Start the server
	e.Logger.Fatal(e.Start(":8000"))
}
