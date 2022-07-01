package module

import (
	"github.com/igorralexsander/stores-manager/internal/domain/services"
	"github.com/igorralexsander/stores-manager/internal/infra/rest/routes"
)

type RestModule struct {
	storeRoute *routes.Store
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
