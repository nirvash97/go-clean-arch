package middleware

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("[%s] %s %s \n", time.Now().UTC().Format(timeFormat), c.Request().Method, c.Request().RequestURI)
		return next(c)
	}
}
