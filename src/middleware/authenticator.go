package middleware

import (
	"github.com/labstack/echo"
	"net/http"
)

func Authenticator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		id := context.Request().Header.Get("id")
		password := context.Request().Header.Get("password")

		// debug
		if id == "user" && password == "password" {
			return next(context)
		}

		return echo.NewHTTPError(http.StatusUnauthorized)
	}
}
