package repository

import (
	"booking/internal/domain/dto"
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
	"context"

	"github.com/gocraft/dbr/v2"
)

type BookingRepository struct {
	*DBConn
}

func NewBookingRepository(db *DBConn) *BookingRepository {
	return &BookingRepository{db}
}

func (r *BookingRepository) GetBookingByID(ctx context.Context, id int64) (entity.Booking, error) {
	var booking dbmodel.Booking

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Select("*").
			From("bookings").
			LeftJoin("inventory", "bookings.inventory_id = inventory.id").
			LeftJoin("resorts", "inventory.resort_id = resorts.id").
			Where("bookings.id = ?", id).
			LoadOne(&booking)
	})

	return dto.BookingFromDB(booking), err
}

func (r *BookingRepository) GetBookingsByUserID(ctx context.Context, userID int64) ([]entity.Booking, error) {
	var bookings []dbmodel.Booking

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("bookings").
			LeftJoin("inventory", "bookings.inventory_id = inventory.id").
			LeftJoin("resorts", "inventory.resort_id = resorts.id").
			Where("bookings.user_id = ?", userID).
			Load(&bookings)

		return err
	})

	return dto.BookingsFromDB(bookings), err
}

func (r *BookingRepository) CreateBooking(ctx context.Context, booking entity.Booking) (entity.Booking, error) {
	dbBooking := dto.BookingToDB(booking)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.InsertInto("bookings").
			Returning("id").
			Columns("user_id", "inventory_id", "total_price", "start_date", "end_date").
			Record(&dbBooking).
			Load(&dbBooking.ID)
	})

	return dto.BookingFromDB(dbBooking), err
}
