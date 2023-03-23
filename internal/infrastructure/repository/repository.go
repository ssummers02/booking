package repository

import (
	"context"

	"github.com/ssummers02/booking/internal/domain"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/domain/service"

	"github.com/gocraft/dbr/v2"
)

type Gateway interface {
	DeleteFrom(table string) *dbr.DeleteStmt
	DeleteBySql(query string, value ...interface{}) *dbr.DeleteStmt
	InsertInto(table string) *dbr.InsertStmt
	InsertBySql(query string, value ...interface{}) *dbr.InsertStmt
	Select(column ...string) *dbr.SelectStmt
	SelectBySql(query string, value ...interface{}) *dbr.SelectStmt
	Update(table string) *dbr.UpdateStmt
	UpdateBySql(query string, value ...interface{}) *dbr.UpdateStmt
}

type DBConn struct {
	*dbr.Connection
}

func (r *DBConn) NewTransaction(ctx context.Context) (entity.AbstractTransaction, error) {
	return r.NewSession(nil).BeginTx(ctx, nil)
}

func (r *DBConn) BeginTx(ctx context.Context, f func(tx *dbr.Tx) error) error {
	tx, err := r.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return domain.NewDBErrorWrap(err)
	}

	defer tx.RollbackUnlessCommitted()

	err = f(tx)
	if err != nil {
		return domain.NewDBErrorWrap(err)
	}

	if tx.Commit() != nil {
		return domain.NewDBErrorWrap(err)
	}

	return nil
}

func NewRepository(db *dbr.Connection) *service.Storages {
	base := &DBConn{db}

	return &service.Storages{
		User:      NewUserRepository(base),
		Resort:    NewResortRepository(base),
		Inventory: NewInventoryRepository(base),
		Booking:   NewBookingRepository(base),
	}
}
