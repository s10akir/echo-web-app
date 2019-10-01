package main

import (
	"net/http"

	"github.com/labstack/echo"
)

const PORT = ":8080"

func main() {
	app := echo.New()

	app.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello World!")
	})

	app.Start(PORT)
}
