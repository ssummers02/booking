package dbmodel

import "time"

type Booking struct {
	ID          int64     `db:"id"`
	UserID      int64     `db:"user_id"`
	InventoryID int64     `db:"inventory_id"`
	Inventory   Inventory `db:"inventory"`
	Resort      Resort    `db:"resort"`
	TotalPrice  float64   `db:"total_price"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
}
