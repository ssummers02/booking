package repository

import (
	"context"
	"errors"

	"github.com/ssummers02/booking/internal/domain"
	"github.com/ssummers02/booking/internal/domain/dto"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/infrastructure/dbmodel"

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
		_, err := tx.Select("bookings.*, inventory.type_id, inventory.resort_id,inventory.price,inventory.photo, resorts.name, resorts.city_id, resorts.owner_id, resorts.description, resorts.address").
			From("bookings").
			LeftJoin("inventory", "bookings.inventory_id = inventory.id").
			LeftJoin("resorts", "inventory.resort_id = resorts.id").
			Where("bookings.user_id = ?", userID).
			OrderDesc("bookings.id").
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
			Columns("user_id", "inventory_id", "total_price", "start_time", "end_time").
			Record(&dbBooking).
			Load(&dbBooking.ID)
	})

	return dto.BookingFromDB(dbBooking), err
}

func (r *BookingRepository) GetBookingsByResort(ctx context.Context, resortID int64) ([]entity.Booking, error) {
	var booking []dbmodel.Booking

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("bookings.*, inventory.type_id, inventory.resort_id,inventory.price,inventory.photo, resorts.name, resorts.city_id, resorts.owner_id, resorts.description, resorts.address").
			From("bookings").
			LeftJoin("inventory", "bookings.inventory_id = inventory.id").
			LeftJoin("resorts", "inventory.resort_id = resorts.id").
			Where("resorts.id = ?", resortID).
			Load(&booking)

		return err
	})

	if errors.Is(err, domain.ErrNotFound) {
		return []entity.Booking{}, nil
	}

	return dto.BookingsFromDB(booking), err
}

func (r *BookingRepository) GetBookingsByOwner(ctx context.Context, ownerID int64) ([]entity.Booking, error) {
	var booking []dbmodel.Booking

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("bookings.*, inventory.type_id, inventory.resort_id,inventory.price,inventory.photo, resorts.name, resorts.city_id, resorts.owner_id, resorts.description, resorts.address").
			From("bookings").
			LeftJoin("inventory", "bookings.inventory_id = inventory.id").
			LeftJoin("resorts", "inventory.resort_id = resorts.id").
			Where("resorts.owner_id = ?", ownerID).
			Load(&booking)

		return err
	})

	if errors.Is(err, domain.ErrNotFound) {
		return []entity.Booking{}, nil
	}

	return dto.BookingsFromDB(booking), err
}

func (r *BookingRepository) StatsBookingInventoryByOwner(ctx context.Context, f entity.StatisticFilter, ownerID int64) ([]entity.Stats, error) {
	var stats []dbmodel.Stats

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := r.selectViewsByPeriod(tx, f)

		_, err := stmt.From("bookings").
			LeftJoin("inventory", "bookings.inventory_id = inventory.id").
			LeftJoin("resorts", "inventory.resort_id = resorts.id").
			Where("resorts.owner_id = ?", ownerID).
			Where("bookings.start_time >= ?", f.StartDate).
			Where("bookings.start_time <= ?", f.EndDate).
			Load(&stats)

		return err
	})

	if errors.Is(err, domain.ErrNotFound) {
		return []entity.Stats{}, nil
	}

	return dto.StatsFromDB(stats), err
}
func (r *BookingRepository) StatsBookingInventoryByResorts(ctx context.Context, f entity.StatisticFilter, id int64) ([]entity.Stats, error) {
	var stats []dbmodel.Stats

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := r.selectViewsByPeriod(tx, f)

		_, err := stmt.From("bookings").
			LeftJoin("inventory", "bookings.inventory_id = inventory.id").
			LeftJoin("resorts", "inventory.resort_id = resorts.id").
			Where("resorts.id = ?", id).
			Where("bookings.start_time >= ?", f.StartDate).
			Where("bookings.start_time <= ?", f.EndDate).
			Load(&stats)

		return err
	})

	if errors.Is(err, domain.ErrNotFound) {
		return []entity.Stats{}, nil
	}

	return dto.StatsFromDB(stats), err
}

func (r *BookingRepository) selectViewsByPeriod(tx *dbr.Tx, f entity.StatisticFilter) *dbr.SelectStmt {
	switch f.GroupBy {
	case entity.StatisticGroupByDay:
		return tx.Select("date_trunc('day', bookings.start_time) as date, count(bookings.inventory_id), bookings.inventory_id").
			GroupBy("date_trunc('day', bookings.start_time), bookings.inventory_id")
	case entity.StatisticGroupByMonth:
		return tx.Select("date_trunc('month', bookings.start_time) as date, count(bookings.inventory_id), bookings.inventory_id").
			GroupBy("date_trunc('month', bookings.start_time), bookings.inventory_id")
	case entity.StatisticGroupByYear:
		return tx.Select("date_trunc('year', bookings.start_time) as date, count(bookings.inventory_id), bookings.inventory_id").
			GroupBy("date_trunc('year', bookings.start_time), bookings.inventory_id")

	default:
		panic("unknown period")
	}
}
