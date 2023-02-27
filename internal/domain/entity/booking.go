package entity

import "time"

type Booking struct {
	ID          int64
	UserID      int64
	InventoryID int64
	Inventory   Inventory
	Resort      Resort
	TotalPrice  float64

	StartDate time.Time
	EndDate   time.Time
	Duration  *int64
}
