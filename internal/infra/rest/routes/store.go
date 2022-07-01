package routes

import (
	"github.com/igorralexsander/stores-manager/internal/domain/model"
	"github.com/igorralexsander/stores-manager/internal/domain/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Store struct {
	storeService services.Store
}

func NewStore(storeService services.Store) *Store {
	return &Store{storeService: storeService}
}

func (r *Store) Register(e *echo.Echo) {
	group := e.Group("store")
	group.POST("", r.Create)
	group.PUT("", r.Update)
	group.GET("", r.GetAll)
}

func (r *Store) Create(c echo.Context) error {
	var payload model.Store
	if err := c.Bind(&payload); err != nil {
		return err
	}
	err := r.storeService.Create(payload)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (r *Store) Update(c echo.Context) error {
	return nil
}

func (r *Store) GetAll(c echo.Context) error {
	return nil
}
