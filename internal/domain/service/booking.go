package service

import (
	"context"
	"errors"

	"github.com/ssummers02/booking/internal/domain"
	"github.com/ssummers02/booking/internal/domain/entity"
)

type BookingService struct {
	repo BookingStorage

	ResortsService   *ResortsService
	InventoryService *InventoryService
	UsersService     *UsersService
}

func NewBookingService(repo BookingStorage, resortsService *ResortsService, inventoryService *InventoryService, usersService *UsersService,
) *BookingService {
	return &BookingService{
		repo:             repo,
		ResortsService:   resortsService,
		InventoryService: inventoryService,
		UsersService:     usersService,
	}
}

func (s *BookingService) GetBookingByID(ctx context.Context, id int64) (entity.Booking, error) {
	booking, err := s.repo.GetBookingByID(ctx, id)
	if err != nil {
		return entity.Booking{}, err
	}

	return booking, nil
}

func (s *BookingService) GetBookingsByUserID(ctx context.Context, userID int64) ([]entity.Booking, error) {
	bookings, err := s.repo.GetBookingsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (s *BookingService) CreateBooking(ctx context.Context, booking entity.Booking) (entity.Booking, error) {
	// проверяем что бронь не пересекается с другими бронями
	err := s.ResortsService.CheckReservation(ctx, booking)
	if err != nil {
		return entity.Booking{}, errors.New("reservation is not available")
	}

	inventory, err := s.InventoryService.GetInventoryByID(ctx, booking.InventoryID)
	if err != nil {
		return entity.Booking{}, err
	}

	booking.TotalPrice = float64(inventory.Price) * booking.EndTime.Sub(booking.StartTime).Hours()

	createBooking, err := s.repo.CreateBooking(ctx, booking)
	if err != nil {
		return entity.Booking{}, err
	}

	return s.GetBookingByID(ctx, createBooking.ID)
}

func (s *BookingService) GetBookingByResortID(ctx context.Context, resortID int64) ([]entity.Booking, error) {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return []entity.Booking{}, domain.NewError(domain.ErrCodeAlreadyExists, "user is not authorized")
	}

	resort, err := s.ResortsService.GetResortByID(ctx, resortID)
	if err != nil {
		return []entity.Booking{}, err
	}

	if resort.OwnerID != user.ID {
		return []entity.Booking{}, errors.New("user is not owner")
	}

	booking, err := s.repo.GetBookingsByResort(ctx, resortID)
	if err != nil {
		return []entity.Booking{}, err
	}

	return booking, nil
}
func (s *BookingService) GetBookingByOwner(ctx context.Context) ([]entity.Booking, error) {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return []entity.Booking{}, domain.NewError(domain.ErrCodeNotAuthorized, "user is not authorized")
	}

	booking, err := s.repo.GetBookingsByOwner(ctx, user.ID)
	if err != nil {
		return []entity.Booking{}, err
	}

	users := make([]int64, 0, len(booking))
	for _, b := range booking {
		users = append(users, b.UserID)
	}

	usersByIDs, err := s.UsersService.GetUsersByIDs(ctx, users)
	if err != nil {
		return nil, err
	}

	for i := range booking {
		booking[i].User = usersByIDs[booking[i].UserID]
	}

	return booking, nil
}

func (s *BookingService) StatsBookingInventoryByOwner(ctx context.Context, f entity.StatisticFilter) ([]entity.Stats, error) {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return []entity.Stats{}, domain.NewError(domain.ErrCodeNotAuthorized, "user is not authorized")
	}

	booking, err := s.repo.StatsBookingInventoryByOwner(ctx, f, user.ID)
	if err != nil {
		return []entity.Stats{}, err
	}

	return booking, nil
}
func (s *BookingService) StatsBookingInventoryByResorts(ctx context.Context, f entity.StatisticFilter, id int64) ([]entity.Stats, error) {
	user, ok := ctx.Value("user").(entity.User)
	if !ok {
		return []entity.Stats{}, domain.NewError(domain.ErrCodeNotAuthorized, "user is not authorized")
	}

	resort, err := s.ResortsService.GetResortByID(ctx, id)
	if err != nil {
		return []entity.Stats{}, err
	}

	if resort.OwnerID != user.ID {
		return []entity.Stats{}, errors.New("user is not owner")
	}

	booking, err := s.repo.StatsBookingInventoryByResorts(ctx, f, id)
	if err != nil {
		return []entity.Stats{}, err
	}

	return booking, nil
}
