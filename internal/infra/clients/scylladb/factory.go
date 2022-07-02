package scylladb

import (
	"github.com/gocql/gocql"
	"github.com/igorralexsander/stores-manager/internal/infra/config"
)

func MakeClient(config *config.DatabaseConfig) *Client {
	return newScyllaDbClient(makeClusterConfig(config))
}

func makeClusterConfig(config *config.DatabaseConfig) *gocql.ClusterConfig {
	clusterConfig := gocql.NewCluster(config.Hosts...)
	clusterConfig.ConnectTimeout = config.ConnectTimeout
	clusterConfig.Timeout = config.ReadTimeout
	clusterConfig.NumConns = config.MaxConnections
	clusterConfig.Keyspace = config.KeySpace
	clusterConfig.Port = config.Port
	clusterConfig.Consistency = gocql.One
	clusterConfig.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	clusterConfig.RetryPolicy = &gocql.SimpleRetryPolicy{
		NumRetries: config.Retries,
	}
	clusterConfig.Authenticator = gocql.PasswordAuthenticator{
		Username: config.User,
		Password: config.Password,
	}
	return clusterConfig
}
