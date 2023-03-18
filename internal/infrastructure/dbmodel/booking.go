package dbmodel

import "time"

type Booking struct {
	ID          int64     `db:"id" json:"id,"`
	UserID      int64     `db:"user_id" json:"user_id"`
	InventoryID int64     `db:"inventory_id" json:"inventory_id"`
	Inventory   Inventory `db:"inventory" json:"inventory"`
	Resort      Resort    `db:"resort" json:"resort"`
	TotalPrice  float64   `db:"total_price" json:"total_price"`
	StartTime   time.Time `db:"start_time" json:"start_time"`
	EndTime     time.Time `db:"end_time" json:"end_time"`
}
