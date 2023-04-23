package dto

import (
	"github.com/ssummers02/booking/internal/delivery/api/restmodel"
	"github.com/ssummers02/booking/internal/domain/entity"
	"github.com/ssummers02/booking/internal/infrastructure/dbmodel"
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
		StartTime: r.StartTime,
		EndTime:   r.EndTime,
	}
}

func ImgToDB(e entity.Img) dbmodel.Img {
	return dbmodel.Img{
		InventoryID: e.InventoryID,
		Name:        e.Name,
		Bytes:       e.Bytes,
	}

}

func ImgFromDB(db dbmodel.Img) entity.Img {
	return entity.Img{
		ID:          db.ID,
		InventoryID: db.InventoryID,
		Name:        db.Name,
		Bytes:       db.Bytes,
	}
}

func ImgFromRest(parseID int64, data []byte, name string) entity.Img {
	return entity.Img{
		InventoryID: parseID,
		Name:        name,
		Bytes:       data,
	}
}
