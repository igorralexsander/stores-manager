package routes

import (
	"github.com/igorralexsander/stores-manager/internal/infra/repository_impl"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Health struct {
	dbChecker repository_impl.DBStatus
}

func NewHealth(dbChecker repository_impl.DBStatus) *Health {
	return &Health{dbChecker: dbChecker}
}

func (r *Health) Register(e *echo.Echo) {
	group := e.Group("/health")
	group.GET("/liveness", r.Liveness)
	group.GET("/readiness", r.Readiness)
}

func (r *Health) Liveness(c echo.Context) error {
	result, err := r.dbChecker.IsUp(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (r *Health) Readiness(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
