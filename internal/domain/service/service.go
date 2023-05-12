package service

import (
	"context"

	"github.com/ssummers02/booking/internal/domain/entity"
)

type Service struct {
	entity.Transactioner

	UsersService     UsersService
	ResortsService   ResortsService
	InventoryService InventoryService
	BookingService   BookingService
	CommentService   CommentService
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
	booking := NewBookingService(r.Booking, resort, inventory, user)
	comment := NewCommentService(r.Comment)

	return &Service{
		UsersService:     *user,
		ResortsService:   *resort,
		InventoryService: *inventory,
		BookingService:   *booking,
		CommentService:   *comment,
	}
}

type Storages struct {
	User      UserStorage
	Resort    ResortStorage
	Inventory InventoryStorage
	Booking   BookingStorage
	Comment   CommentStorage
}
