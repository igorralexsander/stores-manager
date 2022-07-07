package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/igorralexsander/stores-manager/internal/data/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/data/entities"
	"github.com/igorralexsander/stores-manager/internal/domain/model"
)

type store struct {
	dbClient *scylladb.Client
}

func NewStoreRepository(dbClient *scylladb.Client) *store {
	return &store{dbClient: dbClient}
}

func (r *store) Save(ctx context.Context, store model.Store) error {
	dbStore := entities.FromDomain(store)
	stm := "INSERT INTO stores(id, name, url, max_timeout) VALUES (?,?,?,?)"
	return r.dbClient.GetSession().ContextQuery(ctx, stm, []string{}).
		Bind(dbStore.ID, dbStore.Name, dbStore.Url, dbStore.MaxTimeout).
		Exec()
}

func (r *store) GetById(ctx context.Context, id uuid.UUID) (*model.Store, error) {
	stm := "SELECT id, name, url, max_timeout FROM stores WHERE id=?"
	var storeDb entities.StoreDB
	err := r.dbClient.GetSession().ContextQuery(ctx, stm, []string{}).
		Bind(scylladb.ToCQLUUID(id)).
		Get(&storeDb)
	if err != nil {
		return nil, err
	}
	return storeDb.MapToDomain(), nil
}

func (r *store) GetByName(ctx context.Context, name string) (*model.Store, error) {
	return nil, nil
}

func (r *store) GetAll(ctx context.Context) (model.Stores, error) {
	stm := "SELECT id, name, url, max_timeout FROM stores"
	var storesDb entities.StoresDB
	err := r.dbClient.GetSession().ContextQuery(ctx, stm, []string{}).
		Select(&storesDb)
	if err != nil {
		return nil, err
	}
	return storesDb.MapToDomain(), nil
}
