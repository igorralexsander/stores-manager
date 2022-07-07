package module

import (
	"github.com/igorralexsander/stores-manager/internal/data/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/data/repository"
	"github.com/igorralexsander/stores-manager/internal/domain/store"
)

type Repository struct {
	scyllaClient     *scylladb.Client
	dbChecker        repository.DBStatus
	storesRepository store.Repository
}

func NewRepoistoryModule(scyllaClient *scylladb.Client) *Repository {
	return &Repository{
		scyllaClient: scyllaClient,
	}
}

func (m *Repository) ProvideDbChecker() repository.DBStatus {
	if m.dbChecker == nil {
		m.dbChecker = repository.NewDbChecker(m.scyllaClient)
	}
	return m.dbChecker
}

func (m *Repository) ProvideStoreRepository() store.Repository {
	if m.storesRepository == nil {
		m.storesRepository = repository.NewStoreRepository(m.scyllaClient)
	}
	return m.storesRepository
}
