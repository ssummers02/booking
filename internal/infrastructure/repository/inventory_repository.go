package repository

import (
	"booking/internal/domain/dto"
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
	"context"

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
