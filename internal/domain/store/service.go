package store

import (
	"context"
	"github.com/google/uuid"
	"github.com/igorralexsander/stores-manager/internal/domain/model"
)

type Service interface {
	Create(ctx context.Context, store model.Store) error
	FindById(ctx context.Context, id uuid.UUID) (*model.Store, error)
	FindAll(ctx context.Context) (model.Stores, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, store model.Store) error {
	if store.ID == nil {
		newId, _ := uuid.NewRandom()
		store.ID = &newId
	}
	return s.repository.Save(ctx, store)
}

func (s *service) FindById(ctx context.Context, id uuid.UUID) (*model.Store, error) {
	store, err := s.repository.GetById(ctx, id)
	return store, err
}

func (s *service) FindAll(ctx context.Context) (model.Stores, error) {
	return s.repository.GetAll(ctx)
}
