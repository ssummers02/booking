package service

import (
	"context"

	"github.com/ssummers02/booking/internal/domain/entity"
)

type UserStorage interface {
	GetUsersByEmail(ctx context.Context, mail string) (entity.User, error)
	CreateUser(ctx context.Context, e entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
	UpdateUser(ctx context.Context, e entity.User) (entity.User, error)
	GetUsersByIDs(ctx context.Context, ids []int64) ([]entity.User, error)
}

type ResortStorage interface {
	GetResorts(ctx context.Context) ([]entity.Resort, error)
	GetResortByID(ctx context.Context, id int64) (entity.Resort, error)
	GetResortByOwnerID(ctx context.Context, id int64) ([]entity.Resort, error)
	CreateResort(ctx context.Context, e entity.Resort) (entity.Resort, error)
	DeleteResort(ctx context.Context, id int64) error
	UpdateResort(ctx context.Context, e entity.Resort) (entity.Resort, error)
	GetResortsByFilter(ctx context.Context, filter entity.Filter) ([]entity.Resort, error)
	GetCities(ctx context.Context) ([]entity.City, error)

	CheckReservation(ctx context.Context, e entity.Booking) error
}

type InventoryStorage interface {
	GetInventoryByID(ctx context.Context, id int64) (entity.Inventory, error)
	GetInventoryByResortID(ctx context.Context, resortID int64) ([]entity.Inventory, error)
	CreateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error)
	DeleteInventory(ctx context.Context, id int64) error
	UpdateInventory(ctx context.Context, e entity.Inventory) (entity.Inventory, error)
	GetInventoriesTypes(ctx context.Context) ([]entity.InventoryType, error)
	GetInventoriesByFilters(ctx context.Context, filter entity.InventoryFilter) ([]entity.Inventory, error)
	UpdateImg(ctx context.Context, e entity.Img) (entity.Img, error)
	GetImgByInventoryID(ctx context.Context, id int64) (entity.Img, error)
}

type BookingStorage interface {
	GetBookingByID(ctx context.Context, id int64) (entity.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID int64) ([]entity.Booking, error)
	CreateBooking(ctx context.Context, booking entity.Booking) (entity.Booking, error)
	GetBookingsByResort(ctx context.Context, resortID int64) ([]entity.Booking, error)
	GetBookingsByOwner(ctx context.Context, ownerID int64) ([]entity.Booking, error)

	StatsBookingInventoryByOwner(ctx context.Context, f entity.StatisticFilter, ownerID int64) ([]entity.Stats, error)
	StatsBookingInventoryByResorts(ctx context.Context, f entity.StatisticFilter, id int64) ([]entity.Stats, error)
}

type CommentStorage interface {
	CreateComment(ctx context.Context, booking entity.Comment) (entity.Comment, error)
	GetCommentByID(ctx context.Context, id int64) (entity.Comment, error)
	GetCommentsByResort(ctx context.Context, id int64) ([]entity.Comment, error)
	GetCommentsByInventory(ctx context.Context, id int64) ([]entity.Comment, error)
	DeleteCommentsByID(ctx context.Context, id int64) error
}
