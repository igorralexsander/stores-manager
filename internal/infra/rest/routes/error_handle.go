package routes

import (
	"github.com/igorralexsander/stores-manager/internal/domain/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleError(c echo.Context, err error) error {
	switch err.(type) {
	case model.DomainError:
		return handleDomainError(c, err.(model.DomainError))
	default:
		return c.NoContent(http.StatusInternalServerError)
	}
}

func handleDomainError(c echo.Context, domainError model.DomainError) error {
	switch domainError.Code {
	case model.AlreadyExists:
		return c.JSON(http.StatusConflict, domainError)
	case model.NotFound:
		return c.JSON(http.StatusNotFound, domainError)
	default:
		return c.JSON(http.StatusInternalServerError, domainError)
	}
}
