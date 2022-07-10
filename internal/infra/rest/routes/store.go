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
	group.DELETE("/:id", r.Delete)
	group.GET("", r.GetAll)
	group.GET("/:id", r.GetById) //important add / or when get param its with /
	group.GET("/group/:group", r.GetByGroup)
}

func (r *Store) Create(c echo.Context) error {
	var payload model.Store
	if err := c.Bind(&payload); err != nil {
		return HandleError(c, err)
	}
	err := r.storeService.Create(c.Request().Context(), payload)
	if err != nil {
		return HandleError(c, err)
	}
	return c.NoContent(http.StatusCreated)
}

func (r *Store) Update(c echo.Context) error {
	return nil
}

func (r *Store) GetAll(c echo.Context) error {
	stores, err := r.storeService.FindAll(c.Request().Context())
	if err != nil {
		return HandleError(c, err)
	}
	return c.JSON(http.StatusOK, stores)
}

func (r *Store) GetById(c echo.Context) error {
	id := c.Param("id")
	storeId, err := uuid.Parse(id)
	if err != nil {
		return HandleError(c, err)
	}
	resultStore, err := r.storeService.FindById(c.Request().Context(), storeId)
	if err != nil {
		return HandleError(c, err)
	}
	return c.JSON(http.StatusOK, resultStore)
}

func (r *Store) Delete(c echo.Context) error {
	id := c.Param("id")
	storeId, err := uuid.Parse(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := r.storeService.Delete(c.Request().Context(), storeId); err != nil {
		return HandleError(c, err)
	}
	return c.NoContent(http.StatusNoContent)

}

func (r *Store) GetByGroup(c echo.Context) error {
	group := c.Param("group")
	if group == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	stores, err := r.storeService.FindByGroup(c.Request().Context(), group)
	if err != nil {
		return HandleError(c, err)
	}
	return c.JSON(http.StatusOK, stores)
}
