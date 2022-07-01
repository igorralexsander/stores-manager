package routes

import "github.com/labstack/echo/v4"

type Base interface {
	Register(e *echo.Echo)
}
