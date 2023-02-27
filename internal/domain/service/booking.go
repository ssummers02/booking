package service

import (
	"booking/internal/domain/entity"
	"context"
	"errors"
)

type BookingService struct {
	repo BookingStorage

	ResortsService   *ResortsService
	InventoryService *InventoryService
}

func NewBookingService(repo BookingStorage, resortsService *ResortsService, inventoryService *InventoryService) *BookingService {
	return &BookingService{
		repo:             repo,
		ResortsService:   resortsService,
		InventoryService: inventoryService,
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

	booking.TotalPrice = float64(inventory.Price) * booking.EndDate.Sub(booking.StartDate).Hours() / 24

	return s.repo.CreateBooking(ctx, booking)
}
