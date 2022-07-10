package store

import (
	"context"
	"github.com/google/uuid"
	"github.com/igorralexsander/stores-manager/internal/domain/model"
)

type Repository interface {
	GetById(ctx context.Context, id uuid.UUID) (*model.Store, error)
	GetByName(ctx context.Context, name string) (*model.Store, error)
	GetAll(ctx context.Context) (model.Stores, error)
	GetByGroup(ctx context.Context, groupName string) (model.Stores, error)
	Save(ctx context.Context, store model.Store) error
	Delete(ctx context.Context, id uuid.UUID) error
}
