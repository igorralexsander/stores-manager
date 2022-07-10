package module

import (
	"github.com/igorralexsander/stores-manager/internal/data/repository"
	"github.com/igorralexsander/stores-manager/internal/domain/store"
	"github.com/igorralexsander/stores-manager/internal/infra/rest/routes"
)

type Rest struct {
	storeRoute  *routes.Store
	healthRoute *routes.Health
}

func NewRestModule() *Rest {
	return &Rest{}
}

func (m *Rest) ProvideStoreRoute(service store.Service) *routes.Store {
	if m.storeRoute == nil {
		m.storeRoute = routes.NewStore(service)
	}
	return m.storeRoute
}

func (m *Rest) ProvideHealthRoute(dbChecker repository.DBStatus) *routes.Health {
	if m.healthRoute == nil {
		m.healthRoute = routes.NewHealth(dbChecker)
	}
	return m.healthRoute
}
