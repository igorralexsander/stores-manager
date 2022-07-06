package module

import (
	"github.com/igorralexsander/stores-manager/internal/infra/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/infra/repository_impl"
)

type RepositoryModule struct {
	dbChecker repository_impl.DBStatus
}

func NewRepoistoryModule() *RepositoryModule {
	return &RepositoryModule{}
}

func (m *RepositoryModule) ProvideDbChecker(scyllaClient *scylladb.Client) repository_impl.DBStatus {
	if m.dbChecker == nil {
		m.dbChecker = repository_impl.NewDbChecker(scyllaClient)
	}
	return m.dbChecker
}
