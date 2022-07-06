package module

import (
	"github.com/igorralexsander/stores-manager/internal/data/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/infra/config"
)

type Clients struct {
	scyllaClient *scylladb.Client
}

func NewClientsModule() *Clients {
	return &Clients{}
}

func (m *Clients) ProvideScyllaClient(config *config.DatabaseConfig) *scylladb.Client {
	if m.scyllaClient == nil {
		m.scyllaClient = scylladb.MakeClient(config)
	}
	return m.scyllaClient
}
