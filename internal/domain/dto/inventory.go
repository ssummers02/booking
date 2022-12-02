package dto

import (
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
)

func InventoryFromDB(db dbmodel.Inventory) entity.Inventory {
	return entity.Inventory{
		ID:       db.ID,
		TypeID:   db.TypeID,
		ResortID: db.ResortID,
		Price:    db.Price,
		Photo:    db.Photo,
	}
}

func InventorysFromDB(db []dbmodel.Inventory) []entity.Inventory {
	inventories := make([]entity.Inventory, len(db))
	for i, v := range db {
		inventories[i] = InventoryFromDB(v)
	}

	return inventories
}

func InventoryToDB(e entity.Inventory) dbmodel.Inventory {
	return dbmodel.Inventory{
		ID:       e.ID,
		TypeID:   e.TypeID,
		ResortID: e.ResortID,
		Price:    e.Price,
		Photo:    e.Photo,
	}
}
