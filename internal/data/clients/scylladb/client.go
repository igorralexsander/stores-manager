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
	instance := &Client{clusterConfig: clusterConfig}
	return instance
}

func (db *Client) Connect() error {
	if db.session == nil {
		if scyllaSession, err := gocqlx.WrapSession(gocql.NewSession(*db.clusterConfig)); err != nil {
			return err
		} else {
			db.session = &scyllaSession
			return nil
		}
	}
	return nil
}

func (db *Client) GetSession() *gocqlx.Session {
	if db.session == nil || db.session.Closed() {
		db.Connect()
	}
	return db.session
}
