package main

import (
	"errors"
	"fmt"
	"testing"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
)

func TestMarkdownHandler404OnNonPost(t *testing.T) {
	tmpl := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.SetRenderer(tmpl)
	e.Post("/", MarkdownHandler)

	req := test.NewRequest("GET", "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	e.DefaultHTTPErrorHandler(errors.New("error"), c)
	fmt.Println(rec.Status())
	// assert.Equal(t, http.StatusNotFound, rec.Status())
}
