package service

import (
	"booking/internal/domain/entity"
	"context"
)

type UserStorage interface {
	entity.Transactioner

	GetUsersByEmail(ctx context.Context, mail string) (entity.User, error)
	CreateUser(ctx context.Context, e entity.User) (entity.User, error)
}

type ResortStorage interface {
	entity.Transactioner

	GetResorts(ctx context.Context) ([]entity.Resort, error)
	GetResortByID(ctx context.Context, id int64) (entity.Resort, error)
	CreateResort(ctx context.Context, e entity.Resort) (entity.Resort, error)
	DeleteResort(ctx context.Context, id int64) error
	UpdateResort(ctx context.Context, e entity.Resort) (entity.Resort, error)
	GetResortsByCityID(ctx context.Context, cityID int64) ([]entity.Resort, error)
	GetCities(ctx context.Context) ([]entity.City, error)
}

type InventoryStorage interface {
	entity.Transactioner

	GetInventoryByID(ctx context.Context, id int64) (entity.Inventory, error)
	GetInventoryByResortID(ctx context.Context, resortID int64) ([]entity.Inventory, error)
	CreateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error)
	DeleteInventory(ctx context.Context, id int64) error
	UpdateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error)
	GetInventoriesTypes(ctx context.Context) ([]entity.InventoryType, error)
}
