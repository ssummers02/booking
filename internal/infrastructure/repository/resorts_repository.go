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

func (r *ResortRepository) GetResortsByCityID(ctx context.Context, cityID int64) ([]entity.Resort, error) {
	var resorts []dbmodel.Resort

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("resorts").
			Where("city_id = ?", cityID).
			Load(&resorts)

		return err
	})

	return dto.ResortsFromDB(resorts), err
}
