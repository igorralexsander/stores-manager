package module

import (
	"github.com/igorralexsander/stores-manager/internal/data/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/data/repository"
)

type Repository struct {
	scyllaClient *scylladb.Client
	dbChecker    repository.DBStatus
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
