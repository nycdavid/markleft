package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func TestMarkdownHandler405OnNonPost(t *testing.T) {
	tmpl := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.SetRenderer(tmpl)
	req, _ := http.NewRequest(echo.GET, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()

	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	MarkdownHandler(c)

	if rec.Code != 405 {
		t.Errorf("Expected status code to be 405, but got %v", rec.Code)
	}
}
