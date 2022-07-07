package routes

import (
	"github.com/google/uuid"
	"github.com/igorralexsander/stores-manager/internal/domain/model"
	"github.com/igorralexsander/stores-manager/internal/domain/store"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Store struct {
	storeService store.Service
}

func NewStore(storeService store.Service) *Store {
	return &Store{storeService: storeService}
}

func (r *Store) Register(e *echo.Echo) {
	group := e.Group("stores")
	group.POST("", r.Create)
	group.PUT("", r.Update)
	group.GET("", r.GetAll)
	group.GET("/:id", r.GetById) //important add / or when get param its with /
}

func (r *Store) Create(c echo.Context) error {
	var payload model.Store
	if err := c.Bind(&payload); err != nil {
		return err
	}
	err := r.storeService.Create(c.Request().Context(), payload)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (r *Store) Update(c echo.Context) error {
	return nil
}

func (r *Store) GetAll(c echo.Context) error {
	stores, err := r.storeService.FindAll(c.Request().Context())
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, stores)
}

func (r *Store) GetById(c echo.Context) error {
	id := c.Param("id")
	storeId, err := uuid.Parse(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	resultStore, err := r.storeService.FindById(c.Request().Context(), storeId)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, resultStore)
}
