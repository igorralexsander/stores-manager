package module

import "github.com/igorralexsander/stores-manager/internal/domain/services"

type Service struct {
	storeService services.Store
}

func NewServiceModule() *Service {
	return &Service{}
}

func (m *Service) ProvideStoreService() services.Store {
	if m.storeService == nil {
		m.storeService = services.NewStoreService()
	}
	return m.storeService
}
