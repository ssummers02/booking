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
	UsersService   UsersService
	ResortsService ResortsService
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
	return &Service{
		UsersService:   *NewUsersService(r.User),
		ResortsService: *NewResortsService(r.Resort),
	}
}

type Storages struct {
	User   UserStorage
	Resort ResortStorage
}
