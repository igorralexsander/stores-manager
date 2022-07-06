package module

import (
	"github.com/igorralexsander/stores-manager/internal/infra/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/infra/config"
)

type ClientsModule struct {
	scyllaClient *scylladb.Client
}

func NewClientsModule() *ClientsModule {
	return &ClientsModule{}
}

func (m *ClientsModule) ProvideScyllaClient(config *config.DatabaseConfig) *scylladb.Client {
	if m.scyllaClient == nil {
		m.scyllaClient = scylladb.MakeClient(config)
	}
	return m.scyllaClient
}
