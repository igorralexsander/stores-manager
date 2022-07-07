package entities

import (
	"github.com/gocql/gocql"
	"github.com/igorralexsander/stores-manager/internal/data/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/domain/model"
	"time"
)

type StoreDB struct {
	ID         gocql.UUID
	Name       string
	Url        string
	MaxTimeout time.Duration
}

func (m StoreDB) MapToDomain() *model.Store {
	return &model.Store{
		ID:         scylladb.FromCQLUUID(m.ID),
		Name:       m.Name,
		Url:        m.Url,
		MaxTimeout: m.MaxTimeout,
	}
}

func FromDomain(domain model.Store) StoreDB {
	return StoreDB{
		ID:         scylladb.ToCQLUUID(*domain.ID),
		Name:       domain.Name,
		Url:        domain.Url,
		MaxTimeout: domain.MaxTimeout,
	}
}

type StoresDB []StoreDB

func (m StoresDB) MapToDomain() model.Stores {
	domainStores := make(model.Stores, 0)
	for _, dbStore := range m {
		domainStores = append(domainStores, *dbStore.MapToDomain())
	}
	return domainStores
}
