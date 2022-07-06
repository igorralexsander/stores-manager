package module

import (
	"github.com/igorralexsander/stores-manager/internal/domain/services"
	"github.com/igorralexsander/stores-manager/internal/infra/repository_impl"
	"github.com/igorralexsander/stores-manager/internal/infra/rest/routes"
)

type RestModule struct {
	storeRoute  *routes.Store
	healthroute *routes.Health
}

func NewRestModule() *RestModule {
	return &RestModule{}
}

func (m *RestModule) ProvideStoreRoute(service services.Store) *routes.Store {
	if m.storeRoute == nil {
		m.storeRoute = routes.NewStore(service)
	}
	return m.storeRoute
}

func (m *RestModule) ProvideHealthRoute(dbChecker repository_impl.DBStatus) *routes.Health {
	if m.healthroute == nil {
		m.healthroute = routes.NewHealth(dbChecker)
	}
	return m.healthroute
}
