package internal

import (
	serviceModule "github.com/igorralexsander/stores-manager/internal/domain/module"
	restModule "github.com/igorralexsander/stores-manager/internal/infra/module"
	"github.com/igorralexsander/stores-manager/internal/infra/rest/routes"
	"github.com/labstack/echo/v4"
)

type App struct {
	serviceModule *serviceModule.ServiceModule
	restModule    *restModule.RestModule
	routes        []routes.Base
}

func NewApplication(serviceModule *serviceModule.ServiceModule, restModule *restModule.RestModule) *App {
	instance := App{
		serviceModule: serviceModule,
		restModule:    restModule,
	}
	instance.routes = instance.createRoutes()
	return &instance
}

func (a *App) createRoutes() []routes.Base {
	apiRoutes := make([]routes.Base, 0)
	apiRoutes = append(apiRoutes, a.provideStoreRoute())
	return apiRoutes
}

func (a *App) provideStoreRoute() *routes.Store {
	storeService := a.serviceModule.ProvideStoreService()
	return routes.NewStore(storeService)
}

func (a *App) RegisterRoutes(e *echo.Echo) {
	for _, route := range a.routes {
		route.Register(e)
	}
}
