package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"text/template"

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

type ReqBody struct {
	Markdown string `json:"markdown"`
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.SetRenderer(t)
	e.Post("/", MarkdownHandler)
	e.Run(standard.New(os.Getenv("PORT")))
}

func MarkdownHandler(ctx echo.Context) error {
	var rb ReqBody
	d := json.NewDecoder(ctx.Request().Body())
	d.Decode(&rb)
	if rb.Markdown == "" {
		return ctx.JSON(http.StatusBadRequest, "Error: Missing markdown key/string")
	}
	output := blackfriday.MarkdownCommon([]byte(rb.Markdown))
	return ctx.Render(http.StatusOK, "markdown-tmpl", string(output))
}

// 1. e.SetRenderer expects an arg that is of type Renderer
// 2. Renderer is an interface that is expected to have a method `Render`
// 3. The Template struct statisfies this requirement because we declare Render as a member.
