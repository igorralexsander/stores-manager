package module

import "github.com/igorralexsander/stores-manager/internal/domain/services"

type ServiceModule struct {
	storeService services.Store
}

func NewServiceModule() *ServiceModule {
	return &ServiceModule{}
}

func (m *ServiceModule) ProvideStoreService() services.Store {
	if m.storeService == nil {
		m.storeService = services.NewStoreService()
	}
	return m.storeService
}
