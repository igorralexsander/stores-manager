package internal

import (
	dataModule "github.com/igorralexsander/stores-manager/internal/data/module"
	serviceModule "github.com/igorralexsander/stores-manager/internal/domain/module"
	infraModule "github.com/igorralexsander/stores-manager/internal/infra/module"
	"github.com/igorralexsander/stores-manager/internal/infra/rest/routes"
	"github.com/labstack/echo/v4"
)

type App struct {
	serviceModule    *serviceModule.Service
	restModule       *infraModule.Rest
	clientsModule    *dataModule.Clients
	repositoryModule *dataModule.Repository
	routes           []routes.Base
}

func NewApplication(serviceModule *serviceModule.Service, restModule *infraModule.Rest, clientsModule *dataModule.Clients, repositoryModule *dataModule.Repository) *App {
	instance := App{
		serviceModule:    serviceModule,
		restModule:       restModule,
		clientsModule:    clientsModule,
		repositoryModule: repositoryModule,
	}
	instance.routes = instance.createRoutes()
	return &instance
}

func (a *App) createRoutes() []routes.Base {
	apiRoutes := make([]routes.Base, 0)
	apiRoutes = append(apiRoutes, a.provideStoreRoute())
	apiRoutes = append(apiRoutes, a.provideHealthRoute())
	return apiRoutes
}

func (a *App) provideStoreRoute() *routes.Store {
	storeService := a.serviceModule.ProvideStoreService()
	return routes.NewStore(storeService)
}

func (a *App) provideHealthRoute() *routes.Health {
	dbChecker := a.repositoryModule.ProvideDbChecker()
	return routes.NewHealth(dbChecker)
}

func (a *App) RegisterRoutes(e *echo.Echo) {
	for _, route := range a.routes {
		route.Register(e)
	}
}
