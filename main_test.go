package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/russross/blackfriday"
)

func TestMarkdownKeyMissingReturns400(t *testing.T) {
	markdownJson := `{"foo": "bar"}`
	tmpl := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.SetRenderer(tmpl)
	req, _ := http.NewRequest(echo.POST, "/", strings.NewReader(markdownJson))
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	MarkdownHandler(c)

	if rec.Code != 400 {
		t.Errorf("Expected status code to be 405, but got %v", rec.Code)
	}
}

func TestMarkdownKeyMissingReturnsError(t *testing.T) {
	errorMsg := "Error: Missing markdown key/string"
	var jsonStr string
	markdownJson := `{"foo": "bar"}`
	tmpl := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.SetRenderer(tmpl)
	req, _ := http.NewRequest(echo.POST, "/", strings.NewReader(markdownJson))
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	MarkdownHandler(c)
	jsonBuf, _ := ioutil.ReadAll(rec.Body)
	json.Unmarshal(jsonBuf, &jsonStr)

	if jsonStr != errorMsg {
		t.Errorf("Expected response to be %s, but got %s", errorMsg, jsonStr)
	}
}

func TestValidMarkdownReturnsMarkup(t *testing.T) {
	markdownJson := `{"markdown": "* Bar"}`
	markdownMarkup := blackfriday.MarkdownCommon([]byte("* Bar"))
	tmpl := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.SetRenderer(tmpl)
	req, _ := http.NewRequest(echo.POST, "/", strings.NewReader(markdownJson))
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	MarkdownHandler(c)

	respStr, _ := ioutil.ReadAll(rec.Body)
	if string(respStr) != string(markdownMarkup) {
		t.Errorf("Expected response to be %s, but got %s", string(markdownMarkup), string(respStr))
	}
}
