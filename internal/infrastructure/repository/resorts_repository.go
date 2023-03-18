package repository

import (
	"booking/internal/domain/dto"
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
	"context"

	"github.com/gocraft/dbr/v2"
)

type ResortRepository struct {
	*DBConn
}

func NewResortRepository(db *DBConn) *ResortRepository {
	return &ResortRepository{db}
}

func (r *ResortRepository) GetResorts(ctx context.Context) ([]entity.Resort, error) {
	var resorts []dbmodel.Resort

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("resorts").
			Load(&resorts)

		return err
	})

	return dto.ResortsFromDB(resorts), err
}

func (r *ResortRepository) GetResortByID(ctx context.Context, id int64) (entity.Resort, error) {
	var resort dbmodel.Resort

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Select("*").
			From("resorts").
			Where("id = ?", id).
			LoadOne(&resort)
	})

	return dto.ResortFromDB(resort), err
}
func (r *ResortRepository) GetResortByOwnerID(ctx context.Context, id int64) ([]entity.Resort, error) {
	var resort []dbmodel.Resort

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("resorts").
			Where("owner_id = ?", id).
			Load(&resort)
		return err
	})

	return dto.ResortsFromDB(resort), err
}

func (r *ResortRepository) CreateResort(ctx context.Context, e entity.Resort) (entity.Resort, error) {
	resort := dto.ResortToDB(e)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.InsertInto("resorts").
			Returning("id").
			Columns("name", "city_id", "description", "address").
			Record(&resort).
			Load(&resort.ID)
	})

	return dto.ResortFromDB(resort), err
}

func (r *ResortRepository) UpdateResort(ctx context.Context, e entity.Resort) (entity.Resort, error) {
	resort := dto.ResortToDB(e)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Update("resorts").
			Set("name", resort.Name).
			Set("city_id", resort.CityID).
			Set("description", resort.Description).
			Set("address", resort.Address).
			Where("id = ?", resort.ID).
			Exec()

		return err
	})

	return dto.ResortFromDB(resort), err
}

func (r *ResortRepository) DeleteResort(ctx context.Context, id int64) error {
	return r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.DeleteFrom("resorts").
			Where("id = ?", id).
			Exec()

		return err
	})
}

func (r *ResortRepository) GetCities(ctx context.Context) ([]entity.City, error) {
	var cities []dbmodel.City

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("cities").
			Load(&cities)

		return err
	})

	return dto.CitiesFromDB(cities), err
}

func (r *ResortRepository) GetResortsByFilter(ctx context.Context, filter entity.Filter) ([]entity.Resort, error) {
	var resorts []dbmodel.Resort

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := tx.Select("resorts.*").
			From("resorts").
			Join("inventory", "resorts.id = inventory.resort_id").
			GroupBy("resorts.id")
		setFilter(stmt, filter)

		_, err := stmt.Load(&resorts)

		return err
	})

	return dto.ResortsFromDB(resorts), err
}

func setFilter(stmt *dbr.SelectStmt, filter entity.Filter) {
	if filter.CityID != nil {
		stmt.Where("city_id = ?", filter.CityID)
	}

	if filter.TypeID != nil {
		stmt.Where("inventory.type_id = ?", filter.TypeID)
	}

	if filter.StartTime != nil && filter.EndTime != nil {
		stmt.LeftJoin("bookings", dbr.And(
			dbr.Expr("bookings.inventory_id = inventory.id"),
		))
		stmt.Where(
			dbr.Or(
				dbr.Expr("bookings.id IS NULL"),
				dbr.Expr("bookings.start_time > ?", filter.EndTime.Format("2006-01-02 15:04:05")),
				dbr.Expr("bookings.end_time < ?", filter.StartTime.Format("2006-01-02 15:04:05")),
			))
		stmt.Having("COUNT(inventory.id) > 0")
	} else {
		stmt.Where("1 = 1") // Placeholder condition to prevent SQL syntax errors
	}
}

// check that the reservation is not included in other reservations.
func (r *ResortRepository) CheckReservation(ctx context.Context, e entity.Booking) error {
	return r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("bookings.*").
			From("bookings").
			Where("inventory_id = ?", e.InventoryID).
			Where("start_time < ?", e.EndTime.Format("2006-01-02 15:04:05")).
			Where("end_time > ?", e.StartTime.Format("2006-01-02 15:04:05")).
			Load(&dbmodel.Booking{})

		return err
	})
}
