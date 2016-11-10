package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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
	newStr := "* Foo \\n* Bar"
	strings.Replace(newStr, "\\n", "\n", -1)
	fmt.Println(newStr)
	header := string(ctx.Request().Header().Get("X-Markdown"))
	markdown := []byte(header)
	output := string(blackfriday.MarkdownCommon(markdown))
	return ctx.Render(http.StatusOK, "markdown-tmpl", output)
}

// 1. e.SetRenderer expects an arg that is of type Renderer
// 2. Renderer is an interface that is expected to have a method `Render`
// 3. The Template struct statisfies this requirement because we declare Render as a member.
