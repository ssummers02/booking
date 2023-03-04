package dto

import (
	"booking/internal/delivery/api/restmodel"
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

func InventoryToRest(e entity.Inventory) restmodel.Inventory {
	return restmodel.Inventory{
		ID:       e.ID,
		TypeID:   e.TypeID,
		ResortID: e.ResortID,
		Price:    e.Price,
		Photo:    e.Photo,
	}
}

func InventorysToRest(e []entity.Inventory) []restmodel.Inventory {
	inventories := make([]restmodel.Inventory, len(e))

	for i, v := range e {
		inventories[i] = InventoryToRest(v)
	}

	return inventories
}

func InventoryFromRest(r restmodel.Inventory) entity.Inventory {
	return entity.Inventory{
		ID:       r.ID,
		TypeID:   r.TypeID,
		ResortID: r.ResortID,
		Price:    r.Price,
		Photo:    r.Photo,
	}
}

func InventoryTypeFromDB(db dbmodel.InventoryType) entity.InventoryType {
	return entity.InventoryType{
		ID:   db.ID,
		Name: db.Name,
	}
}

func InventoryTypesFromDB(db []dbmodel.InventoryType) []entity.InventoryType {
	inventoryTypes := make([]entity.InventoryType, len(db))
	for i, v := range db {
		inventoryTypes[i] = InventoryTypeFromDB(v)
	}

	return inventoryTypes
}

func InventoryTypeToRest(e entity.InventoryType) restmodel.InventoryType {
	return restmodel.InventoryType{
		ID:   e.ID,
		Name: e.Name,
	}
}

func InventoryTypesToRest(e []entity.InventoryType) []restmodel.InventoryType {
	inventoryTypes := make([]restmodel.InventoryType, len(e))

	for i, v := range e {
		inventoryTypes[i] = InventoryTypeToRest(v)
	}

	return inventoryTypes
}

func InventoryFilterFromRest(r restmodel.InventoryFilter) entity.InventoryFilter {
	return entity.InventoryFilter{
		ResortID:  r.ResortID,
		TypeID:    r.TypeID,
		StartDate: r.StartDate,
		Duration:  r.Duration,
	}
}
