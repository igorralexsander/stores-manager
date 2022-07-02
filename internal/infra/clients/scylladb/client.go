package scylladb

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Client struct {
	clusterConfig *gocql.ClusterConfig
	session       *gocqlx.Session
}

func newScyllaDbClient(clusterConfig *gocql.ClusterConfig) *Client {
	return &Client{clusterConfig: clusterConfig}
}

func (db *Client) connect(clusterConfig gocql.ClusterConfig) error {
	if db.session == nil {
		if scyllaSession, err := gocqlx.WrapSession(gocql.NewSession(clusterConfig)); err != nil {
			return err
		} else {
			db.session = &scyllaSession
			return nil
		}
	}
	return nil
}

func (db *Client) GetSession() *gocqlx.Session {
	return db.session
}
