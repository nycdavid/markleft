package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	e.Get("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})
	e.Run(standard.New(":1323"))
}
