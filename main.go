package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/russross/blackfriday"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.SetRenderer(t)
	e.Get("/", Hello)
	e.Run(standard.New(":1323"))
}

func Hello(ctx echo.Context) error {
	markdown := []byte(ctx.Request().Header().Get("X-Markdown"))
	output := blackfriday.MarkdownBasic(markdown)
	return ctx.Render(http.StatusOK, "markdown-tmpl", string(output))
}

// 1. e.SetRenderer expects an arg that is of type Renderer
// 2. Renderer is an interface that is expected to have a method `Render`
// 3. The Template struct statisfies this requirement because we declare Render as a member.
