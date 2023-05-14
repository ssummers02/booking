package repository

import (
	"context"

	"github.com/ssummers02/booking/internal/domain/dto"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/infrastructure/dbmodel"

	"github.com/gocraft/dbr/v2"
)

type InventoryRepository struct {
	*DBConn
}

func NewInventoryRepository(db *DBConn) *InventoryRepository {
	return &InventoryRepository{db}
}

func (r *InventoryRepository) GetInventoryByID(ctx context.Context, id int64) (entity.Inventory, error) {
	var inventory dbmodel.Inventory

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Select("*").
			From("inventory").
			Where("id = ?", id).
			LoadOne(&inventory)
	})

	return dto.InventoryFromDB(inventory), err
}

func (r *InventoryRepository) GetInventoryByResortID(ctx context.Context, resortID int64) ([]entity.Inventory, error) {
	var inventories []dbmodel.Inventory

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("inventory").
			Where("resort_id = ?", resortID).
			Load(&inventories)

		return err
	})

	return dto.InventorysFromDB(inventories), err
}

func (r *InventoryRepository) CreateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error) {
	inventory := dto.InventoryToDB(e)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.InsertInto("inventory").
			Returning("id").
			Columns("type_id", "resort_id", "price", "photo").
			Record(&inventory).
			Load(&inventory.ID)
	})

	return dto.InventoryFromDB(inventory), err
}

func (r *InventoryRepository) UpdateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error) {
	inventory := dto.InventoryToDB(e)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Update("inventory").
			Set("price", inventory.Price).
			Set("photo", inventory.Photo).
			Where("id = ?", inventory.ID).
			Exec()
		if err != nil {
			return err
		}

		return err
	})

	return dto.InventoryFromDB(inventory), err
}

func (r *InventoryRepository) DeleteInventory(ctx context.Context, id int64) error {
	return r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.DeleteFrom("inventory").
			Where("id = ?", id).
			Exec()

		return err
	})
}

func (r *InventoryRepository) GetInventoriesTypes(ctx context.Context) ([]entity.InventoryType, error) {
	var types []dbmodel.InventoryType

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Select("*").
			From("inventory_type").
			Load(&types)

		return err
	})

	return dto.InventoryTypesFromDB(types), err
}
func (r *InventoryRepository) GetInventoriesByFilters(ctx context.Context, filter entity.InventoryFilter) ([]entity.Inventory, error) {
	var inventories []dbmodel.Inventory

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		stmt := tx.Select("inventory.*").
			From("inventory").
			Join("resorts", "resorts.id = inventory.resort_id")

		setInventoryFilter(stmt, filter)

		_, err := stmt.Load(&inventories)

		return err
	})

	return dto.InventorysFromDB(inventories), err
}

func setInventoryFilter(stmt *dbr.SelectStmt, filter entity.InventoryFilter) {
	stmt.Where("resort_id = ?", filter.ResortID)

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

func (r *InventoryRepository) UpdateImg(ctx context.Context, e entity.Img) (entity.Img, error) {
	img := dto.ImgToDB(e)

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.Update("inventory").
			Set("data", img.Bytes).
			Set("name", img.Name).
			Where("id = ?", img.InventoryID).
			Exec()
		if err != nil {
			return err
		}

		return nil
	})

	return dto.ImgFromDB(img), err
}

func (r *InventoryRepository) GetImgByInventoryID(ctx context.Context, id int64) (entity.Img, error) {
	var img dbmodel.Img

	err := r.BeginTx(ctx, func(tx *dbr.Tx) error {
		return tx.Select("*").
			From("inventory").
			Where("id = ?", id).
			LoadOne(&img)
	})

	return dto.ImgFromDB(img), err
}
