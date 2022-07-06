package repository_impl

import (
	"context"
	"github.com/igorralexsander/stores-manager/internal/infra/clients/scylladb"
	"time"
)

type dbChecker struct {
	dbClient *scylladb.Client
}

type Status struct {
	CurrentTimestamp time.Time `json:"currentTimestamp"`
}

type DBStatus interface {
	IsUp(ctx context.Context) (*Status, error)
}

func NewDbChecker(dbClient *scylladb.Client) *dbChecker {
	return &dbChecker{
		dbClient: dbClient,
	}
}

func (db *dbChecker) IsUp(ctx context.Context) (*Status, error) {
	stm := "SELECT dateOf(now()) FROM system.local;"
	query := db.dbClient.GetSession().ContextQuery(ctx, stm, []string{})
	var result time.Time
	if err := query.Get(&result); err != nil {
		return nil, err
	}
	return &Status{CurrentTimestamp: result}, nil
}
