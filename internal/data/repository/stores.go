package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/igorralexsander/stores-manager/internal/data/clients/scylladb"
	"github.com/igorralexsander/stores-manager/internal/data/entities"
	"github.com/igorralexsander/stores-manager/internal/domain/model"
)

const (
	selectCommand = "SELECT id, name, url, max_timeout, group FROM stores"
	insertCommand = "INSERT INTO stores(id, name, url, max_timeout, group) VALUES (?,?,?,?,?) IF NOT EXISTS"
)

type store struct {
	dbClient *scylladb.Client
}

func NewStoreRepository(dbClient *scylladb.Client) *store {
	return &store{dbClient: dbClient}
}

func (r *store) Save(ctx context.Context, store model.Store) error {
	dbStore := entities.FromDomain(store)
	applied, err := r.dbClient.GetSession().ContextQuery(ctx, insertCommand, []string{}).
		Bind(dbStore.ID, dbStore.Name, dbStore.Url, dbStore.MaxTimeout, dbStore.Group).
		ExecCASRelease()
	if err != nil {
		return err
	}
	if !applied {
		return model.AlreadyExistsError(store.ID)
	}
	return nil
}

func (r *store) GetById(ctx context.Context, id uuid.UUID) (*model.Store, error) {
	stm := fmt.Sprintf("%s WHERE id=?", selectCommand)
	var storeDb entities.StoreDB
	_, err := r.dbClient.GetSession().ContextQuery(ctx, stm, []string{}).
		Bind(scylladb.ToCQLUUID(id)).
		GetCASRelease(&storeDb)
	if err != nil {
		if err.Error() == gocql.ErrNotFound.Error() {
			return nil, model.NotFoundError(id)
		}
		return nil, err
	}
	return storeDb.MapToDomain(), nil
}

func (r *store) GetByName(ctx context.Context, name string) (*model.Store, error) {
	return nil, nil
}

func (r *store) GetAll(ctx context.Context) (model.Stores, error) {
	var storesDb entities.StoresDB
	err := r.dbClient.GetSession().ContextQuery(ctx, selectCommand, []string{}).
		SelectRelease(&storesDb)
	if err != nil {
		return nil, err
	}
	return storesDb.MapToDomain(), nil
}

func (r *store) GetByGroup(ctx context.Context, groupName string) (model.Stores, error) {
	stm := fmt.Sprintf("%s WHERE group=?", selectCommand)
	var storesDb entities.StoresDB
	err := r.dbClient.GetSession().
		ContextQuery(ctx, stm, []string{}).
		Bind(groupName).
		SelectRelease(&storesDb)
	if err != nil {
		return nil, err
	}
	return storesDb.MapToDomain(), nil
}

func (r *store) Delete(ctx context.Context, id uuid.UUID) error {
	stm := "DELETE FROM stores WHERE id=?;"
	return r.dbClient.GetSession().
		ContextQuery(ctx, stm, []string{}).
		Bind(scylladb.ToCQLUUID(id)).
		ExecRelease()
}
