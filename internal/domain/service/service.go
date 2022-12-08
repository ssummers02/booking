package service

import (
	"booking/internal/domain/entity"
	"context"
)

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Fatal(args ...interface{})
	Sync() error
}

type Service struct {
	entity.Transactioner
	UsersService     UsersService
	ResortsService   ResortsService
	InventoryService InventoryService
}

func (s *Service) DoTransaction(ctx context.Context, f func() error) (err error) {
	tx, err := s.NewTransaction(ctx)
	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	err = f()

	if err != nil {
		return err
	}

	return tx.Commit()
}

func NewServices(r *Storages) *Service {
	user := NewUsersService(r.User)
	resort := NewResortsService(r.Resort)
	inventory := NewInventoryService(r.Inventory, resort)
	return &Service{
		UsersService:     *user,
		ResortsService:   *resort,
		InventoryService: *inventory,
	}
}

type Storages struct {
	User      UserStorage
	Resort    ResortStorage
	Inventory InventoryStorage
}
