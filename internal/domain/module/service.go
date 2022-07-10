package module

import (
	"github.com/igorralexsander/stores-manager/internal/domain/store"
)

type Service struct {
	storeService store.Service
}

func NewServiceModule() *Service {
	return &Service{}
}

func (m *Service) ProvideStoreService(repository store.Repository) store.Service {
	if m.storeService == nil {
		m.storeService = store.NewService(repository)
	}
	return m.storeService
}
