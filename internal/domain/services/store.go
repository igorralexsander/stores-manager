package services

import (
	"github.com/google/uuid"
	"github.com/igorralexsander/stores-manager/internal/domain/model"
)

type storeService struct {
	repository Repository
}

type Store interface {
	Create(store model.Store) error
}

type Repository interface {
	Save(store model.Store) error
	GetById(id uuid.UUID) (model.Store, error)
	GetByName(name string) (model.Store, error)
}

func NewStoreService(repository Repository) *storeService {
	return &storeService{
		repository: repository,
	}
}

func (s *storeService) Create(store model.Store) error {

	return nil
}
